package vm

import (
	"fmt"
	"io"

	"encoding/binary"
)

var enc = binary.LittleEndian

const (
	ExpNone = iota
	ExpHalt
	ExpAddr
	ExpPC
)

// the virtual machine
type C struct {
	gprs []uint32
	pc   uint32
	mem  []byte
	exp  int

	tsc uint64
}

func New(memSize int) *C {
	if memSize%1024 != 0 {
		panic("memory not aligned to 1024")
	}
	if memSize == 0 {
		panic("zero memory")
	}

	ret := new(C)
	ret.gprs = make([]uint32, 16)
	ret.mem = make([]byte, memSize)

	return ret
}

func (c *C) chkpc() bool {
	a := c.pc
	if a&0x3 != 0 {
		c.expaddr()
		return false
	}
	if a+4 > c.memlen() {
		c.expaddr()
		return false
	}
	return true
}

func (c *C) Run(start uint32) int {
	c.pc = start
	c.tsc = 0
	return c.Resume()
}

func (c *C) Resume() int {
	c.exp = ExpNone
	for c.exp == ExpNone {
		c.step()
	}

	return c.exp
}

func (c *C) Step() int {
	c.exp = ExpNone
	return c.step()
}

func (c *C) step() int {
	if !c.chkpc() {
		return c.exp
	}

	inst := c.rdw(c.pc)
	c.pc += 4

	if (inst >> 28) != 0 {
		c.i28(inst)
	} else if (inst >> 24) != 0 {
		c.i24(inst)
	} else if (inst >> 16) != 0 {
		c.i16(inst)
	} else if (inst >> 12) != 0 {
		c.i12(inst)
	} else if (inst >> 8) != 0 {
		c.i8(inst)
	} else if (inst >> 4) != 0 {
		c.i4(inst)
	} else {
		c.i0(inst)
	}

	c.tsc++

	return c.exp
}

func (c *C) i28(i uint32) {
	if (i >> 30) == 1 {
		c.j(i)
	}
	if (i >> 30) == 2 {
		c.jal(i)
	}
	if (i >> 28) == 1 {
		c.syscall(i & 0x0fffffff)
	}
}

func (c *C) i24(inst uint32) {
	x := uint8(inst & 0xf)
	y := uint8((inst >> 4) & 0xf)
	i := uint16(inst >> 8)
	code := uint8((inst >> 24) & 0xf)
	switch code {
	case 1:
		c.addi(x, y, i)
	case 2:
		c.andi(x, y, i)
	case 3:
		c.ori(x, y, i)
	case 4:
		c.slti(x, y, i)
	case 5:
		c.lw(x, y, i)
	case 6:
		c.lh(x, y, i)
	case 7:
		c.lhu(x, y, i)
	case 8:
		c.lb(x, y, i)
	case 9:
		c.lbu(x, y, i)
	case 10:
		c.sw(x, y, i)
	case 11:
		c.sh(x, y, i)
	case 12:
		c.sb(x, y, i)
	case 13:
		c.beq(x, y, i)
	case 14:
		c.bne(x, y, i)
	}
}

func (c *C) i20(inst uint32) {
	x := uint8(inst & 0xf)
	i := uint16(inst >> 4)
	code := uint8((inst >> 24) & 0xf)
	switch code {
	case 1:
		c.lui(x, i)
	}
}

func (c *C) i16(inst uint32) {
	x := uint8(inst & 0xf)
	y := uint8((inst >> 4) & 0xf)
	p := uint8((inst >> 8) & 0xf)
	q := uint8((inst >> 12) & 0xf)
	code := uint8((inst >> 16) & 0xf)

	switch code {
	case 1:
		c.mul(x, y, p, q)
	case 2:
		c.mulu(x, y, p, q)
	case 3:
		c.div(x, y, p, q)
	case 4:
		c.divu(x, y, p, q)
	}
}

func (c *C) i12(inst uint32) {
	x := uint8(inst & 0xf)
	y := uint8((inst >> 4) & 0xf)
	z := uint8((inst >> 8) & 0xf)
	code := uint8((inst >> 12) & 0xf)

	switch code {
	case 1:
		c.add(x, y, z)
	case 2:
		c.sub(x, y, z)
	case 3:
		c.and(x, y, z)
	case 4:
		c.or(x, y, z)
	case 5:
		c.xor(x, y, z)
	case 6:
		c.nor(x, y, z)
	case 7:
		c.slt(x, y, z)
	case 8:
		c.sll(x, y, z)
	case 9:
		c.srl(x, y, z)
	case 10:
		c.sra(x, y, z)
	case 11:
		c.sllv(x, y, z)
	case 12:
		c.srlv(x, y, z)
	case 13:
		c.srav(x, y, z)
	}
}

func (c *C) i8(inst uint32) {}

func (c *C) i4(inst uint32) {
	x := uint8(inst & 0xf)
	code := uint8((inst >> 4) & 0xf)
	switch code {
	case 1:
		c.jr(x)
	}
}

func (c *C) i0(inst uint32) {
	code := uint8(inst & 0xf)
	switch code {
	case 0:
		c.halt()
	}
}

func (c *C) w(i uint8, v uint32) {
	if i == 0 {
		return
	}
	c.gprs[i] = v
}
func (c *C) r(i uint8) uint32 {
	if i == 0 {
		return 0
	}
	return c.gprs[i]
}
func (c *C) rs(i uint8) int32    { return int32(c.r(i)) }
func (c *C) ws(i uint8, v int32) { c.w(i, uint32(v)) }
func bin(c bool) uint32 {
	if c {
		return 1
	}
	return 0
}
func (c *C) trap(t int)      { c.exp = t }
func (c *C) expaddr() uint32 { c.trap(ExpAddr); return 0 }
func (c *C) memlen() uint32  { return uint32(len(c.mem)) }

// i0: 0000 000c
func (c *C) halt() { c.trap(ExpHalt) }

// i4: 0000 00cx
func (c *C) jr(x uint8) { c.pc = c.r(x) }

// i12: 0000 czyx
func (c *C) add(x, y, z uint8)  { c.w(x, c.r(y)+c.r(z)) }
func (c *C) sub(x, y, z uint8)  { c.w(x, c.r(y)-c.r(z)) }
func (c *C) and(x, y, z uint8)  { c.w(x, c.r(y)&c.r(z)) }
func (c *C) or(x, y, z uint8)   { c.w(x, c.r(y)|c.r(z)) }
func (c *C) xor(x, y, z uint8)  { c.w(x, c.r(y)^c.r(z)) }
func (c *C) nor(x, y, z uint8)  { c.w(x, ^(c.r(y) | c.r(z))) }
func (c *C) slt(x, y, z uint8)  { c.w(x, bin(c.r(y) < c.r(z))) }
func (c *C) sll(x, y, z uint8)  { c.w(x, c.r(y)<<z) }
func (c *C) srl(x, y, z uint8)  { c.w(x, c.r(y)>>z) }
func (c *C) sra(x, y, z uint8)  { c.ws(x, c.rs(y)>>z) }
func (c *C) sllv(x, y, z uint8) { c.w(x, c.r(y)<<c.r(z)) }
func (c *C) srlv(x, y, z uint8) { c.w(x, c.r(y)>>c.r(z)) }
func (c *C) srav(x, y, z uint8) { c.ws(x, c.rs(y)>>c.r(z)) }

// i16: 000c pqyx
func (c *C) mul(x, y, p, q uint8) {
	_p, _q := c.rs(p), c.rs(q)
	m := uint64(int64(_p) * int64(_q))
	c.ws(x, int32(m>>32))
	c.ws(y, int32(m&0xffffffff))
}
func (c *C) mulu(x, y, p, q uint8) {
	_p, _q := c.r(p), c.r(q)
	m := uint64(_p) * uint64(_q)
	c.w(x, uint32(m>>32))
	c.w(y, uint32(m&0xffffffff))
}
func (c *C) div(x, y, p, q uint8) {
	_p, _q := c.rs(p), c.rs(q)
	if _q == 0 {
		c.ws(x, 0)
		c.ws(y, 0)
		return
	}
	c.ws(x, _p/_q)
	c.ws(y, _p%_q)
}

func (c *C) divu(x, y, p, q uint8) {
	_p, _q := c.r(p), c.r(q)
	if _q == 0 {
		c.ws(x, 0)
		c.ws(y, 0)
		return
	}
	c.w(x, _p/_q)
	c.w(y, _p%_q)
}

// i20: 00ci iiix

func se(c uint16) int32  { return int32(int16(c)) }
func ze(c uint16) uint32 { return uint32(c) }
func seb(c uint8) int32  { return int32(int8(c)) }
func zeb(c uint8) uint32 { return uint32(c) }

func (c *C) lui(x uint8, i uint16) {
	c.w(x, c.r(x)&0x0000ffff+(ze(i)<<16))
}

// i24: 0cii iiyx

func (c *C) addr(y uint8, i uint16) uint32 {
	return uint32(c.rs(y) + se(i))
}
func (c *C) adw(y uint8, i uint16) uint32 {
	a := c.addr(y, i)
	if a&0x3 != 0 {
		return c.expaddr()
	}
	if a+4 > c.memlen() {
		return c.expaddr()
	}
	return a
}
func (c *C) adhw(y uint8, i uint16) uint32 {
	a := c.addr(y, i)
	if a&0x1 != 0 {
		return c.expaddr()
	}
	if a+2 > c.memlen() {
		return c.expaddr()
	}
	return a
}

func (c *C) adb(y uint8, i uint16) uint32 {
	a := c.addr(y, i)
	if a >= c.memlen() {
		return c.expaddr()
	}
	return a
}

func (c *C) rdw(a uint32) uint32  { return enc.Uint32(c.mem[a : a+4]) }
func (c *C) rdhw(a uint32) uint16 { return enc.Uint16(c.mem[a : a+2]) }
func (c *C) rdb(a uint32) uint8   { return c.mem[a] }

func (c *C) wrw(a uint32, v uint32)  { enc.PutUint32(c.mem[a:a+4], v) }
func (c *C) wrhw(a uint32, v uint16) { enc.PutUint16(c.mem[a:a+4], v) }
func (c *C) wrb(a uint32, v uint8)   { c.mem[a] = v }

func (c *C) addi(x, y uint8, i uint16) { c.ws(x, c.rs(y)+se(i)) }
func (c *C) andi(x, y uint8, i uint16) { c.w(x, c.r(y)&ze(i)) }
func (c *C) ori(x, y uint8, i uint16)  { c.w(x, c.r(y)|ze(i)) }
func (c *C) slti(x, y uint8, i uint16) { c.w(x, bin(c.rs(y) < se(i))) }

func (c *C) lw(x, y uint8, i uint16)  { c.w(x, c.rdw(c.adw(y, i))) }
func (c *C) lh(x, y uint8, i uint16)  { c.ws(x, se(c.rdhw(c.adhw(y, i)))) }
func (c *C) lhu(x, y uint8, i uint16) { c.w(x, ze(c.rdhw(c.adhw(y, i)))) }
func (c *C) lb(x, y uint8, i uint16)  { c.ws(x, seb(c.rdb(c.adb(y, i)))) }
func (c *C) lbu(x, y uint8, i uint16) { c.w(x, zeb(c.rdb(c.adb(y, i)))) }

func (c *C) sw(x, y uint8, i uint16) { c.wrw(c.adw(y, i), c.r(x)) }
func (c *C) sh(x, y uint8, i uint16) { c.wrhw(c.adhw(y, i), uint16(c.r(x))) }
func (c *C) sb(x, y uint8, i uint16) { c.wrb(c.adb(y, i), uint8(c.r(x))) }

func (c *C) beq(x, y uint8, i uint16) {
	if c.r(x) == c.r(y) {
		c.pc += uint32(se(i)) << 2
	}
}
func (c *C) bne(x, y uint8, i uint16) {
	if c.r(x) != c.r(y) {
		c.pc += uint32(se(i)) << 2
	}
}

// i28: xiii iiii

// x = 01ii
func (c *C) j(i uint32) { c.pc = i << 2 }

// x = 10ii
func (c *C) jal(i uint32) { c.w(15, c.pc); c.j(i) }

// x = 0001
func (c *C) syscall(i uint32) {
	/* todo */
}

func (c *C) Load(m []byte, offset uint32) {
	n := uint32(len(m))
	copy(c.mem[offset:offset+n], m[:n])
}

func (c *C) PrintRegs(out io.Writer) {
	fmt.Fprintf(out, "pc=%08x", c.pc)

	for i := uint8(0); i < 16; i++ {
		if i%4 == 0 {
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, " ")
		}
		fmt.Fprintf(out, "$%x=%08x", i, c.r(i))
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "tsc=%d\n", c.tsc)
}

func (c *C) ClearTSC() { c.tsc = 0 }
func (c *C) TSC() uint64 { return c.tsc }

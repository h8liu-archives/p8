package vm

import (
	"encoding/binary"
)

var enc = binary.LittleEndian

const (
	AddressError = iota
)

type Core struct {
	gprs   []uint32
	hi, lo uint32
	pc     uint32
	mem    []byte
}

func (self *Core) NewCore(memSize int) *Core {
	if memSize%1024 != 0 {
		panic("memory not aligned to 1024")
	}
	if memSize == 0 {
		panic("zero memory")
	}

	ret := new(Core)
	ret.gprs = make([]uint32, 32)
	ret.mem = make([]byte, memSize)

	return ret
}

func (self *Core) w(i int, v uint32) {
	if i == 0 {
		return
	}
	self.gprs[i] = v
}

func (self *Core) r(i int) uint32 {
	if i == 0 {
		return 0
	}
	return self.gprs[i]
}

func (self *Core) trap(t uint32) { panic("todo") }

func (c *Core) Addu(d, s, t int) { c.w(d, c.r(s)+c.r(t)) }
func (c *Core) Subu(d, s, t int) { c.w(d, c.r(s)-c.r(t)) }

func (c *Core) Addiu(t, s int, C uint16) { c.w(t, c.r(s)+signExt(C)) }

func (c *Core) Mult(s, t int) {
	r := uint64(int64(c.r(s)) * int64(c.r(t)))
	c.lo = uint32(r)
	c.hi = uint32(r >> 32)
}

func (c *Core) Multu(s, t int) {
	r := uint64(c.r(s)) * uint64(c.r(t))
	c.lo = uint32(r)
	c.hi = uint32(r >> 32)
}

func (c *Core) Div(s, t int) {
	_s := int32(c.r(s))
	_t := int32(c.r(t))
	if _t == 0 {
		c.lo = 0
		c.hi = 0
		return
	}
	c.lo = uint32(_s / _t)
	c.hi = uint32(_s % _t)
}

func (c *Core) Divu(s, t int) {
	_s := c.r(s)
	_t := c.r(t)
	if _t == 0 {
		c.lo = 0
		c.hi = 0
		return
	}
	c.lo = _s / _t
	c.hi = _s % _t
}

func (c *Core) word(addr uint32) uint32 {
	return enc.Uint32(c.mem[addr : addr+4])
}

func (c *Core) halfWord(addr uint32) uint16 {
	return enc.Uint16(c.mem[addr : addr+4])
}

func (c *Core) writeWord(addr uint32, v uint32) {
	enc.PutUint32(c.mem[addr:addr+4], v)
}

func (c *Core) writeHalfWord(addr uint32, v uint16) {
	enc.PutUint16(c.mem[addr:addr+2], v)
}

func (c *Core) writeByte(addr uint32, v uint8) {
	c.mem[addr] = v
}

func (c *Core) checkWordAddr(addr uint32) bool {
	return (addr&0x2) == 0 && addr+4 <= c.memLen()
}

func (c *Core) checkHalfWordAddr(addr uint32) bool {
	return (addr&0x1) == 0 && addr+2 <= c.memLen()
}

func (c *Core) checkAddr(addr uint32) bool {
	return addr < c.memLen()
}

func (c *Core) memLen() uint32 { return uint32(len(c.mem)) }

func signExt(C uint16) uint32 { return uint32(int32(int16(C))) }

func (c *Core) Lw(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkWordAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.w(t, c.word(addr))
}

func (c *Core) Lh(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkHalfWordAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.w(t, uint32(int32(int16(c.halfWord(addr)))))
}

func (c *Core) Lhu(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkHalfWordAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.w(t, uint32(c.halfWord(addr)))
}

func (c *Core) Lb(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.w(t, uint32(int32(int8(c.mem[addr]))))
}

func (c *Core) Lbu(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.w(t, uint32(c.mem[addr]))
}

func (c *Core) Sw(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkWordAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.writeWord(addr, c.r(t))
}

func (c *Core) Sh(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkHalfWordAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.writeHalfWord(addr, uint16(c.r(t)))
}

func (c *Core) Sb(t, s int, C uint16) {
	addr := c.r(s) + signExt(C)
	if !c.checkAddr(addr) {
		c.trap(AddressError)
		return
	}

	c.writeByte(addr, uint8(c.r(t)))
}

func (c *Core) Mfhi(d int) { c.w(d, c.hi) }
func (c *Core) Mflo(d int) { c.w(d, c.lo) }

func (c *Core) And(d, s, t int)         { c.w(d, c.r(s)&c.r(t)) }
func (c *Core) Andi(t, s int, C uint16) { c.w(t, c.r(s)&uint32(C)) }
func (c *Core) Or(d, s, t int)          { c.w(d, c.r(s)|c.r(t)) }
func (c *Core) Ori(t, s int, C uint16)  { c.w(t, c.r(s)|uint32(C)) }
func (c *Core) Xor(d, s, t int)         { c.w(d, c.r(s)^c.r(d)) }
func (c *Core) Nor(d, s, t int)         { c.w(d, ^(c.r(s) | c.r(d))) }

func bin(c bool) uint32 {
	if c {
		return 1
	}
	return 0
}

func (c *Core) Slt(d, s, t int) { c.w(d, bin(int32(c.r(s)) < int32(c.r(t)))) }
func (c *Core) Slti(t, s int, C uint16) {
	c.w(t, bin(int32(c.r(s)) < int32(signExt(C))))
}

func (c *Core) Sll(d, t, s int)  { c.w(d, c.r(t)<<uint(s)) }
func (c *Core) Srl(d, t, s int)  { c.w(d, c.r(t)>>uint(s)) }
func (c *Core) Sra(d, t, s int)  { c.w(d, uint32(int32(c.r(t))>>uint(s))) }
func (c *Core) Sllv(d, t, s int) { c.w(d, c.r(t)<<c.r(s)) }
func (c *Core) Srlv(d, t, s int) { c.w(d, c.r(t)>>c.r(s)) }
func (c *Core) Srav(d, t, s int) { c.w(d, uint32(int32(c.r(t))>>c.r(s))) }

func (c *Core) relPC(C uint16) uint32 {
	return c.pc + uint32(4*int32(signExt(C)))
}

func (c *Core) Beq(s, t int, C uint16) {
	if c.r(s) == c.r(t) {
		c.pc = c.relPC(C)
	}
}

func (c *Core) Bne(s, t int, C uint16) {
	if c.r(s) != c.r(t) {
		c.pc = c.relPC(C)
	}
}

func (c *Core) J(C uint32)   { c.pc = (c.pc & 0xfc000000) + (C & 0x03ffffff) }
func (c *Core) Jr(s int)     { c.pc = c.r(s) }
func (c *Core) Jal(C uint32) { c.w(31, c.pc); c.J(C) }

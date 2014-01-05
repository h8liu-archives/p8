package vm

import (
	"fmt"
	"io"
)

// the virtual machine
type C struct {
	gprs []uint64
	pc   uint64

	exp int
	tsc uint64

	mem []byte
}

func New(memSize int) *C {
	if memSize%1024 != 0 {
		panic("memory not aligned to 1K")
	}
	if memSize == 0 {
		panic("zero memory")
	}

	ret := new(C)
	ret.gprs = make([]uint64, 16)
	ret.mem = make([]byte, memSize)

	return ret
}

func (c *C) step() int {
	// read
	i := c.rdd(c.pc)
	c.pc += 8

	// exec
	c.inst(i)

	// clean up
	c.gprs[0] = 0
	u64p(c.mem[0:8], 0)
	c.tsc++

	return c.exp
}

func (c *C) Run(start uint64) int {
	c.pc = start
	c.tsc = 0
	return c.Resume()
}

func (c *C) Resume() int {
	c.exp = ExcepNone
	for c.exp == ExcepNone {
		c.step()
	}

	return c.exp
}

func (c *C) Step() int {
	c.exp = ExcepNone
	return c.step()
}

func (c *C) Load(m []byte, offset uint64) {
	if offset%8 != 0 {
		panic("offset not aligned")
	}
	n := uint64(len(m))
	copy(c.mem[offset:offset+n], m[:n])
}

func (c *C) PrintRegs(out io.Writer) {
	fmt.Fprintf(out, "pc=%016x", c.pc)
	r := c.gprs

	for i := uint8(0); i < 16; i++ {
		if i%4 == 0 {
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, " ")
		}
		fmt.Fprintf(out, "$%x=%016x", i, r[i])
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "tsc=%d\n", c.tsc)
}

func (c *C) ClearTSC()   { c.tsc = 0 }
func (c *C) TSC() uint64 { return c.tsc }

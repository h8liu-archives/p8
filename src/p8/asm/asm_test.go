package asm

import (
	"testing"
)

type progTest struct {
	prog *Prog
	res  *resultCheck
}

type resultCheck struct {
	regs map[uint8]uint64
}

func (rc *resultCheck) R(i uint8, v uint64) {
	rc.regs[i] = v
}

func p1() *Prog {
	p := NewProg()

	f := p.F("main")
	f.Addi(1, 0, 5)
	f.Add(2, 0, 0)
	f.L("loop")
	f.Add(2, 2, 1)
	f.Addi(1, 1, -1)
	f.Bne(0, 1, "loop")
	f.Halt()

	return p
}

func p2() *Prog {
	p := NewProg()

	f := p.F("main")
	f.Jal("f")
	f.Halt()

	f = p.F("f")
	f.Jr(15)

	return p
}

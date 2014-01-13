package asm

import (
	"p8/vm"
	"testing"
)

func TestProgs(t *testing.T) {
	o := func(p *Prog, ttl uint64) *vm.VM {
		buf := p.Assemble(vm.PageStart(1))
		page := vm.NewPage(vm.PermAll)
		copy(page.Bytes, buf)

		v := vm.New(page)
		v.TTL = ttl

		v.ResumeAt(vm.PageStart(1))
		return v
	}

	reg := func(vm *vm.VM, r uint8, v uint64) {
		got := vm.R(r)
		if got != v {
			t.Errorf("r%d expect %016x, got %016x", r, v, got)
		}
	}

	// program: sum 1 to 10
	p := NewProg()
	f := p.F("main")
	f.Addi(1, 0, 10)
	f.Add(2, 0, 0)
	f.L("loop")
	f.Add(2, 2, 1)
	f.Addi(1, 1, -1)
	f.Bne(0, 1, "loop")
	f.Halt()
	v := o(p, 100)
	reg(v, 2, 55)

	// program: call another function
	p = NewProg()
	f = p.F("main")
	f.Addi(1, 0, 27)
	f.Jal("f")
	f.Addi(2, 0, 22)
	f.Halt()
	f = p.F("f")
	f.Addi(1, 0, 30)
	f.Jr(15)

	v = o(p, 100)
	reg(v, 1, 30)
	reg(v, 2, 22)
}

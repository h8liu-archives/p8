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
		v.MapPage(2, vm.NewPage(vm.PermWrite))

		v.TTL = ttl

		r := v.ResumeAt(vm.PageStart(1))
		if r != vm.ExcepHalt {
			t.Errorf("vm exit with %d", r)
		}

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

	// program: fabonacci number, fab(10)
	p = func() *Prog {
		p := NewProg()

		/*
			argument    <- base stack pointer
			return value
			return pc	<- call1 stack pointer
			sum tmp
			argument
			return value
			return pc   <- call2 stack pointer
		*/

		arg := int32(10)
		f := p.F("main")
		f.Addi(14, 0, int32(vm.PageStart(2)))
		f.Addi(1, 0, arg)
		f.Sd(1, 14, 0)     // argument for calling
		f.Addi(14, 14, 16) // prepare stack for calling
		f.Jal("fab")
		f.Ld(1, 14, -8)     // fetch return value
		f.Addi(14, 14, -16) // restore stack from calling
		f.Halt()

		f = p.F("fab")
		f.Sd(15, 14, 0)     // store pc pointer
		f.Ld(1, 14, -16)    // load the argument
		f.Beq(1, 0, "ret1") // fab(0)
		f.Addi(1, 1, -1)    // $2 = arg - 1
		f.Beq(1, 0, "ret1") // fab(1)

		// sub call fab(arg-1)
		f.Sd(1, 14, 16)    // argument for calling
		f.Addi(14, 14, 32) // prepare stack for calling
		f.Jal("fab")
		f.Ld(2, 14, -8)     // fetch return value
		f.Addi(14, 14, -32) // restore stack from calling
		f.Sd(2, 14, 8)      // save return value

		// sub call fab(arg-2)
		f.Ld(1, 14, -16)   // load the argument again
		f.Addi(1, 1, -2)   // $3 = arg - 2
		f.Sd(1, 14, 16)    // argument for calling
		f.Addi(14, 14, 32) // prepare stack for calling
		f.Jal("fab")
		f.Ld(2, 14, -8)     // fetch return value
		f.Addi(14, 14, -32) // retore stack from calling
		f.Ld(1, 14, 8)      // load tmp
		f.Add(1, 1, 2)
		f.Beq(0, 0, "out") // jump to out

		f.L("ret1")
		f.Addi(1, 0, 1) // set 1 to 1

		f.L("out")
		f.Sd(1, 14, -8) // return 1

		f.Ld(15, 14, 0) // restore pc
		f.Jr(15)        // return
		return p
	}()
	v = o(p, 3000)
	reg(v, 1, 89)
}

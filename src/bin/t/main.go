package main

import (
	"fmt"
	"io"
	"os"

	"p8/asm"
	"p8/vm"
)

func FprintRegs(out io.Writer, vm *vm.VM) {
	fmt.Fprintf(out, "pc=%016x", vm.PC())

	for i := uint8(0); i < 16; i++ {
		if i%4 == 0 {
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, " ")
		}
		fmt.Fprintf(out, "$%x=%016x", i, vm.R(i))
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "tsc=%d ttl=%d\n", vm.TSC, vm.TTL)
}

func PrintRegs(vm *vm.VM) { FprintRegs(os.Stdout, vm) }

func p0() *asm.Prog {
	p := asm.NewProg()

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
	f.Sd(1, 14, 0) // argument for calling
	f.Addi(14, 14, 16) // prepare stack for calling
	f.Jal("fab")
	f.Ld(1, 14, -8) // fetch return value
	f.Addi(14, 14, -16) // restore stack from calling
	f.Halt()

	f = p.F("fab")
	f.Sd(15, 14, 0) // store pc pointer
	f.Ld(1, 14, -16) // load the argument
	f.Beq(1, 0, "ret1") // fab(0)
	f.Addi(1, 1, -1) // $2 = arg - 1
	f.Beq(1, 0, "ret1") // fab(1)

	// sub call fab(arg-1)
	f.Sd(1, 14, 16) // argument for calling
	f.Addi(14, 14, 32) // prepare stack for calling
	f.Jal("fab")
	f.Ld(2, 14, -8) // fetch return value
	f.Addi(14, 14, -32) // restore stack from calling
	f.Sd(2, 14, 8) // save return value

	// sub call fab(arg-2)
	f.Ld(1, 14, -16) // load the argument again
	f.Addi(1, 1, -2) // $3 = arg - 2
	f.Sd(1, 14, 16) // argument for calling
	f.Addi(14, 14, 32) // prepare stack for calling
	f.Jal("fab")
	f.Ld(2, 14, -8) // fetch return value
	f.Addi(14, 14, -32) // retore stack from calling
	f.Ld(1, 14, 8) // load tmp
	f.Add(1, 1, 2)
	f.Beq(0, 0, "out") // jump to out

	f.L("ret1")
	f.Addi(1, 0, 1) // set 1 to 1

	f.L("out")
	f.Sd(1, 14, -8) // return 1

	f.Ld(15, 14, 0) // restore pc
	f.Jr(15) // return

	return p
}

func main() {
	p := p0()

	buf := p.Assemble(vm.PageStart(1))
	p.Fprint(os.Stdout)

	page := vm.NewPage(vm.PermAll)
	copy(page.Bytes, buf)

	stack := vm.NewPage(vm.PermWrite)
	v := vm.New(page)
	v.MapPage(2, stack) // map a stack

	v.Log = os.Stdout
	v.TTL = 3000
	e := v.ResumeAt(vm.PageStart(1))
	fmt.Printf("e=%d\n", e)
	PrintRegs(v)
}

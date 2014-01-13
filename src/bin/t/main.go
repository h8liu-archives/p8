package main

import (
	"fmt"
	"io"
	"os"

	"p8/asm"
	. "p8/opcode"
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

func p1() *asm.Prog {
	p := asm.NewProg()
	f := p.F("main")
	f.I(OpXYI(Addi, 1, 0, 5))
	f.I(OpXPQ(Add, 2, 0, 0))
	f.L("loop")
	f.I(OpXPQ(Add, 2, 2, 1))
	f.I(OpXYIs(Addi, 1, 1, -1))
	f.I(OpXY(Bne, 0, 1)).L("loop")
	f.I(Op(Halt))
	return p
}

func p2() *asm.Prog {
	p := asm.NewProg()
	f := p.F("main")
	f.I(Op(J | Jal)).L("f")
	f.I(Op(Halt))

	f = p.F("f")
	f.I(OpP(Jr, 15))

	return p
}

func main() {
	p := p2()

	buf := p.Assemble(vm.PageStart(1))
	p.Fprint(os.Stdout)

	page := vm.NewPage(vm.PermAll)
	copy(page.Bytes, buf)

	v := vm.New(page)
	v.Log = os.Stdout
	v.TTL = 100
	e := v.ResumeAt(vm.PageStart(1))
	fmt.Printf("e=%d\n", e)
	PrintRegs(v)
}

package main

import (
	"os"

	"m8/asm"
	"m8/vm"
)

func main() {
	v := vm.New(4096 * 8)                   // 8 pages
	v.Load(asm.AssembleFile("a.asm"), 4096) // load at page 1
	v.Run(4096)
	v.PrintRegs(os.Stdout)
	v.Resume()
	v.PrintRegs(os.Stdout)
}

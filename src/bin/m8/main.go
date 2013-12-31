package main

import (
	"os"

	"m8"
)

func main() {
	c := m8.NewVM(4096 * 8)                // 8 pages
	c.Load(m8.AssembleFile("a.asm"), 4096) // load at page 1
	c.Run(4096)
	c.PrintRegs(os.Stdout)
}

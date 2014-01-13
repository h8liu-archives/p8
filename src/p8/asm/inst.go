package asm

import (
	"fmt"
	"io"

	. "p8/opcode"
)

type Inst struct {
	Line     uint64
	JmpLabel string
	Error    error
}

func newInst(line uint64) *Inst {
	ret := new(Inst)
	ret.Line = line
	return ret
}

func (self *Inst) L(label string) {
	self.JmpLabel = label
}

func (self *Inst) I(line uint64) {
	self.Line = line
}

func (self *Inst) E(e error) {
	self.Error = e
}

func (self *Inst) Fprint(out io.Writer) {
	op := Opcode(self.Line)
	if (op & J) != 0 {
		if (op & Jal) != 0 {
			fmt.Fprintf(out, "jal %s", self.JmpLabel)
		} else {
			fmt.Fprintf(out, "j %s", self.JmpLabel)
		}
	} else {
		switch op {
		case Beq:
			fmt.Fprintf(out, "beq %s", self.JmpLabel)
		case Bne:
			fmt.Fprintf(out, "bne %s", self.JmpLabel)
		default:
			fmt.Fprint(out, InstStr(self.Line))
		}
	}

	if self.Error != nil {
		fmt.Fprintf(out, "; error: %v", self.Error)
	}
}

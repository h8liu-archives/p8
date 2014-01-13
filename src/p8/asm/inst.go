package asm

import (
	"io"
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

func (self *Inst) J(label string) {
	self.JmpLabel = label
}

func (self *Inst) I(line uint64) {
	self.Line = line
}

func (self *Inst) E(e error) {
	self.Error = e
}

func (self *Inst) Fprint(out io.Writer) {
	panic("todo")
}

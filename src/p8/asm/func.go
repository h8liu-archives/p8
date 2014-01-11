package asm

import (
	"bytes"
	"fmt"

	"p8/opcode"
)

type Func struct {
	Name   string
	Pos    uint64
	insts  []*Inst
	labels map[string]*Label
}

func newFunc(name string) *Func {
	ret := new(Func)
	ret.Name = name
	ret.insts = make([]*Inst, 0, 1024)
	ret.labels = make(map[string]*Label)
	return ret
}

func (self *Func) I(i uint64) *Inst {
	ret := newInst(i)
	self.insts = append(self.insts, ret)
	return ret
}

func (self *Func) L(lab string) (*Label, error) {
	ret := newLabel(lab).P(uint64(len(self.insts)))
	_, exists := self.labels[lab]
	if exists {
		return nil, fmt.Errorf("label %s already exists", lab)
	}
	self.labels[lab] = ret
	return ret, nil
}

func (self *Func) locateLabel(label string, prog *Prog) (uint64, error) {
	lab, exists := self.labels[label]
	if exists {
		return self.Pos + lab.Pos, nil
	}

	return prog.funcPos(label)
}

func (self *Func) assembleInto(out *bytes.Buffer, prog *Prog) error {
	for i, inst := range self.insts {
		line := inst.Line
		op := opcode.Opcode(line)

		if op == opcode.Beq || op == opcode.Bne {
			if inst.JmpLabel == "" {
				return fmt.Errorf("inst %d: beq/bne with empty label", i)
			}

			thisPos := self.Pos + (uint64(i) << 3) + 8
			thatPos, e := self.locateLabel(inst.JmpLabel, prog)
			if e != nil {
				return e
			}

			diff64 := int64(thatPos-thisPos) >> 3
			diff := int32(diff64)
			if int64(diff) != diff64 {
				return fmt.Errorf("inst %d: beq/bne out of range", i)
			}

			inst.I(opcode.InstSetI(line, uint32(diff)))
		} else if op == opcode.J || op == opcode.Jal {
			if inst.JmpLabel == "" {
				return fmt.Errorf("inst %d: j/jal with empty label", i)
			}

			pos, e := self.locateLabel(inst.JmpLabel, prog)
			if e != nil {
				return e
			}

			inst.I(opcode.OpJ(op, pos>>3))
		} else {
			if inst.JmpLabel != "" {
				return fmt.Errorf("inst %d: label not empty on non-jump", i)
			}
		}
	}

	return nil
}

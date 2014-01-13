package asm

import (
	"bytes"
	"fmt"
	"io"

	"p8/opcode"
)

type Func struct {
	Name     string
	Pos      uint64
	insts    []*Inst
	labels   []*Label
	labelMap map[string]*Label
}

func newFunc(name string) *Func {
	ret := new(Func)
	ret.Name = name
	ret.insts = make([]*Inst, 0, 1024)
	ret.labels = make([]*Label, 0, 1024)
	ret.labelMap = make(map[string]*Label)
	return ret
}

func (self *Func) I(i uint64) *Inst {
	ret := newInst(i)
	self.insts = append(self.insts, ret)
	return ret
}

func (self *Func) Label(lab string) (*Label, error) {
	ret := newLabel(lab).P(uint64(len(self.insts)))
	_, exists := self.labelMap[lab]
	if exists {
		return nil, fmt.Errorf("label %s already exists", lab)
	}
	self.labelMap[lab] = ret
	self.labels = append(self.labels, ret)
	return ret, nil
}

func (self *Func) L(lab string) *Label {
	ret, e := self.Label(lab)
	if e != nil {
		panic(e)
	}
	return ret
}

func (self *Func) locateLabel(label string, prog *Prog) (uint64, error) {
	lab, exists := self.labelMap[label]
	if exists {
		return self.Pos + lab.Pos, nil
	}

	return prog.funcPos(label)
}

func (self *Func) assembleInto(out *bytes.Buffer, prog *Prog) {
	for i, inst := range self.insts {
		line := inst.Line
		op := opcode.Opcode(line)

		if op == opcode.Beq || op == opcode.Bne {
			if inst.JmpLabel == "" {
				inst.E(fmt.Errorf("empty label"))
				continue
			}

			thisPos := self.Pos + (uint64(i) << 3) + 8
			thatPos, e := self.locateLabel(inst.JmpLabel, prog)
			if e != nil {
				inst.E(e)
				continue
			}

			diff64 := int64(thatPos-thisPos) >> 3
			diff := int32(diff64)
			if int64(diff) != diff64 {
				inst.E(fmt.Errorf("out of range"))
				continue
			}
			inst.I(opcode.InstSetI(line, uint32(diff)))
		} else if op == opcode.J || op == opcode.Jal {
			if inst.JmpLabel == "" {
				inst.E(fmt.Errorf("empty label"))
				continue
			}

			pos, e := self.locateLabel(inst.JmpLabel, prog)
			if e != nil {
				inst.E(e)
				continue
			}
			inst.I(opcode.OpJ(op, pos>>3))
		} else {
			if inst.JmpLabel != "" {
				inst.E(fmt.Errorf("non-empty label"))
			}
		}
	}
}

func (self *Func) Fprint(out io.Writer) {
	fmt.Fprintf(out, "func %s {\n", self.Name)
	labIndex := 0
	for i, inst := range self.insts {
		for labIndex < len(self.labels) &&
			self.labels[labIndex].Pos == uint64(i) {
			fmt.Fprintf(out, "%s:\n", self.labels[labIndex].Name)
			labIndex++
		}

		fmt.Fprintf(out, "    ")
		inst.Fprint(out)
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out, "}")
	fmt.Fprintln(out)
}

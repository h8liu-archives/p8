package asm

import (
	"bytes"
	"fmt"
)

type Prog struct {
	funcs map[string]*Func
}

func NewProg() *Prog {
	ret := new(Prog)
	ret.funcs = make(map[string]*Func)

	return ret
}

func (self *Prog) Func(name string) (*Func, error) {
	_, exists := self.funcs[name]
	if exists {
		return nil, fmt.Errorf("function %s exists", name)
	}

	ret := newFunc(name)
	self.funcs[name] = ret
	return ret, nil
}

func (self *Prog) funcPos(name string) (uint64, error) {
	f, exists := self.funcs[name]
	if !exists {
		return 0, fmt.Errorf("label %s not exists", name)
	}

	return f.Pos, nil
}

func (self *Prog) Assemble(start uint64) ([]byte, error) {
	ret := new(bytes.Buffer)
	p := start

	// assign positions
	for _, f := range self.funcs {
		f.Pos = p
		p += uint64(len(f.insts)) * 8
	}

	for _, f := range self.funcs {
		e := f.assembleInto(ret, self)
		if e != nil {
			return nil, e
		}
	}

	return ret.Bytes(), nil
}

package asm

import (
	"bytes"
	"fmt"
	"io"
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

func (self *Prog) F(name string) *Func {
	ret, e := self.Func(name)
	if e != nil {
		panic(e)
	}
	return ret
}

func (self *Prog) funcPos(name string) (uint64, error) {
	f, exists := self.funcs[name]
	if !exists {
		return 0, fmt.Errorf("label %s not exists", name)
	}

	return f.Pos, nil
}

func (self *Prog) Assemble(start uint64) []byte {
	ret := new(bytes.Buffer)
	p := start

	// assign positions
	for _, f := range self.funcs {
		f.Pos = p
		p += uint64(len(f.insts)) * 8
	}

	for _, f := range self.funcs {
		f.assembleInto(ret, self)
	}

	return ret.Bytes()
}

func (self *Prog) Fprint(out io.Writer) {
	for _, f := range self.funcs {
		f.Fprint(out)
		fmt.Fprintln(out)
	}
}

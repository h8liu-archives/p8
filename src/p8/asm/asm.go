package asm

import (
	"fmt"
	"io"
)

type Assembler struct {
	Name string
	In   io.Reader
	Log  io.Writer

	lineno int
	errored bool
	toks []string
}

func (self *Assembler) errorf(f string, args ...interface{}) {
	if self.errored {
		return
	}

	fmt.Fprintf(self.Log, "%s:%d: ", self.Name, self.lineno)
	fmt.Fprintf(self.Log, f, args...)
	fmt.Fprintln(self.Log)

	self.errored = true
}

func (self *Assembler) addLabel() {
	if len(self.toks) == 0 {
		self.errorf("empty label")
		return
	}

	if len(self.toks) > 1 {
		self.errorf("unexpected ':'")
		return
	}

	// TOOD: check if the label is reserved identifier
	// TODO: dup check, add the label to index
	panic("todo")

	self.clear()
}

func (self *Assembler) addLine() {
	if len(self.toks) == 0 {
		return // empty linie
	}

	panic("todo")

	self.clear()
}

func (self *Assembler) pushToken(tok string) {
	if !isIdentNum(rune(tok[0])) {
		self.errorf("invalid token: %#v", tok)
		return
	}

	if len(self.toks) < cap(self.toks) {
		self.toks = append(self.toks, tok)
	} else {
		self.errorf("too many tokens")
	}
}

func (self *Assembler) clear() {
	self.toks = self.toks[0:0]
}

func (self *Assembler) lineNext() {
	self.lineno++
	self.errored = false
}

func (self *Assembler) token(tok string) {
	if tok[0] == ';' {
		return
	}

	switch tok {
	case "":
		panic("bug")
	case "\n":
		self.addLine()
		self.lineNext()
	case ":":
		self.addLabel()
	default:
		self.pushToken(tok)
	}
}

func (self *Assembler) Assemble() ([]byte, error) {
	self.toks = make([]string, 0, 8)

	self.lineno = 0
	self.lineNext()

	scanner := newScanner(self.In)
	for scanner.Scan() {
		self.token(scanner.Text())
	}

	return nil, scanner.Err()
}

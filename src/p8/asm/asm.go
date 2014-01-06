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
}

func (self *Assembler) errorf(f string, args ...interface{}) {
	fmt.Fprintf(self.Log, "%s:%d: ", self.Name, self.lineno)
	fmt.Fprintf(self.Log, f, args...)
	fmt.Fprintln(self.Log)
}

func (self *Assembler) addLabel(lab string) {
	if lab == "" {
		self.errorf("empty label")
		return
	}

	// TOOD: check if the label is reserved identifier
	// TODO: dup check, add the label to index
	panic("todo")
}

func (self *Assembler) addLine(op string, args []string) {
	panic("todo")
}

func (self *Assembler) Assemble() ([]byte, error) {
	self.lineno = 1
	args := make([]string, 0, 16)
	cur := ""

	scanner := newScanner(self.In)
	for scanner.Scan() {
		tok := scanner.Text()
		if tok[0] == ';' {
			continue
		}

		switch tok {
		case "":
			panic("bug")
		case "\n":
			self.lineno++
			if cur != "" {
				self.addLine(cur, args)
				cur = ""
				args = args[0:0]
			}
		case ":":
			if len(args) > 0 {
				self.errorf("unexpected: %#v", tok)
			} else {
				self.addLabel(cur)
				cur = ""
			}
		default:
			if isIdentNum(rune(tok[0])) {
				if cur == "" {
					cur = tok
				} else {
					if len(args) < cap(args) {
						args = append(args, tok)
					}
				}
			} else {
				self.errorf("invalid token: %#v", tok)
			}
		}
	}

	return nil, scanner.Err()
}

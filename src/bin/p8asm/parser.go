package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"p8/asm"
)

type Parser struct {
	prog     *asm.Prog
	curFunc  *asm.Func
	errors   []error
	curFile  string
	lineno   int
	gotFatal bool
}

func NewParser() *Parser {
	self := new(Parser)
	self.prog = asm.NewProg()
	self.errors = make([]error, 0, 1024)
	return self
}

func (self *Parser) addError(e error) {
	self.errors = append(self.errors, e)
}

func (self *Parser) ioerrorf(f string, a ...interface{}) {
	self.addError(fmt.Errorf(f, a...))
}

func (self *Parser) errorf(f string, a ...interface{}) {
	header := fmt.Sprintf("%s:%d ", self.curFile, self.lineno)
	msg := fmt.Sprintf(f, a...)
	self.addError(errors.New(header + msg))
}

func (self *Parser) fatalf(f string, a ...interface{}) {
	self.errorf(f, a...)
	self.gotFatal = true
}

func (self *Parser) ParseFile(path string) {
	fin, e := os.Open(path)
	if e != nil {
		self.ioerrorf("open file '%s': %v", path, e)
		return
	}

	defer fin.Close()
	self.parse(fin)
}

func trimLine(line string) string {
	line = trim(line)
	// trim comment
	comment := strings.Index(line, ";")
	if comment > 0 {
		line = line[:comment]
	}
	return line
}

func (self *Parser) Parse(in io.Reader, filename string) {
	self.curFile = filename
}

func (self *Parser) parse(in io.Reader) {
	scanner := bufio.NewScanner(in)

	self.lineno = 0
	self.curFunc = nil
	for scanner.Scan() {
		self.lineno++
		line := trimLine(scanner.Text())
		if line == "" {
			continue
		}

		self.parseLine(line)

		if self.gotFatal {
			break
		}
	}

	e := scanner.Err()
	if e != nil {
		self.ioerrorf("scanner error: %v", e)
	}
}

func trim(s string) string { return strings.TrimSpace(s) }

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

func isValidLabel(label string) bool {
	if len(label) == 0 {
		return false
	}

	for i, r := range label[1:] {
		if i == 0 {
			if !isLetter(r) && r != '_' {
				return false
			}
		} else {
			if !isLetter(r) && !isNumber(r) && r != '_' {
				return false
			}
		}
	}
	return true
}

func (self *Parser) parseFunc(line string) {
	assert(strings.HasPrefix(line, "func"))
	label := trim(line[4:])
	if label == "" {
		self.errorf("empty func label")
	} else if !isValidLabel(label) {
		self.errorf("invalid func label")
	} else {
		var e error
		self.curFunc, e = self.prog.Func(label)
		if e != nil {
			// TODO: report previous declare position
			self.fatalf("%v", e)
		}
	}
}

func (self *Parser) parseLabel(line string) {
	// TODO:
}

func (self *Parser) parseInst(line string, key string) {
	assert(strings.HasPrefix(line, key))
	// TODO:
}

func (self *Parser) parseLine(line string) {
	key := line
	firstSpace := strings.IndexAny(key, " \t")
	if firstSpace > 0 {
		key = line[:firstSpace]
	}

	if key == "func" {
		self.parseFunc(line)
	} else {
		if self.curFunc == nil {
			self.errorf("must start with func")
		} else if strings.HasSuffix(line, ":") {
			self.parseLabel(line)
		} else {
			self.parseInst(line, key)
		}
	}
}

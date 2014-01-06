package asm

import (
	"io"
	"bufio"
)

func Assemble(in io.Reader) []byte {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		
	}

	return nil
}


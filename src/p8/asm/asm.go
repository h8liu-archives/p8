package asm

import (
	"fmt"
	"io"
)

func Assemble(in io.Reader) []byte {
	scanner := newScanner(in)

	for scanner.Scan() {
		token := scanner.Text()
		fmt.Printf("%#v\n", token)
	}

	return nil
}

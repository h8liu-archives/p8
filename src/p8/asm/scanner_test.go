package asm

import (
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	tokens := func(s string) []string {
		ret := make([]string, 0, 1024)
		in := strings.NewReader(s)
		scanner := newScanner(in)
		for scanner.Scan() {
			ret = append(ret, scanner.Text())
		}
		if scanner.Err() != nil {
			panic("bug")
		}

		return ret
	}

	c := func(s string, expect []string) {
		got := tokens(s)
		for i := 0; i < len(got) && i < len(expect); i++ {
			if got[i] != expect[i] {
				t.Errorf("token %d, expect %#v, got %#v,",
					i, expect[i], got[i])
			}
		}
		if len(got) != len(expect) {
			t.Errorf("expect %d tokens, got %d tokens",
				len(expect), len(got))
		}
	}

	c(` add r0, r1, r2
		sub r3, r4, r5 ; some comment here

	main:
		xor r3, r4, r7

	loop:slt r7, r0, r3
		jne r0, r1, loop
		halt
	    ; some more comment
	; and that's it`, []string{
		"add", "r0", "r1", "r2", "\n",
		"sub", "r3", "r4", "r5", "; some comment here", "\n",
		"\n",
		"main", ":", "\n",
		"xor", "r3", "r4", "r7", "\n",
		"\n",
		"loop", ":", "slt", "r7", "r0", "r3", "\n",
		"jne", "r0", "r1", "loop", "\n",
		"halt", "\n",
		"; some more comment", "\n",
		"; and that's it",
	})
}

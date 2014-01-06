package asm

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

func isSpace(r rune) bool {
	if r == '\n' {
		return false
	}
	return r == ',' || unicode.IsSpace(r)
}

func isIdentNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func scanToken(data []byte, atEOF bool) (adv int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// skip leading non-endl white spaces
	var r rune
	width := 0
	start := 0
	for ; start < len(data); start += width {
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}

	if start >= len(data) {
		return 0, nil, nil // need more
	}

	end := start + width
	if isIdentNum(r) {
		// start of an identifier or number
		for ; end < len(data); end += width {
			r, width = utf8.DecodeRune(data[end:])
			if !isIdentNum(r) {
				return end, data[start:end], nil
			}
		}

		if end >= len(data) && !atEOF {
			return 0, nil, nil // need more
		}
	} else if r == ';' { // start of comment
		for ; end < len(data); end += width {
			r, width = utf8.DecodeRune(data[end:])
			if r == '\n' {
				return end, data[start:end], nil
			}
		}

		if end >= len(data) && !atEOF {
			return 0, nil, nil // need more
		}
	} else if r == ':' || r == '\n' {
		// special char, don't need to do anything
	} else {
		// invalid char, don't need to do anything
	}

	return end, data[start:end], nil
}

func newScanner(in io.Reader) *bufio.Scanner {
	ret := bufio.NewScanner(in)
	ret.Split(scanToken)
	return ret
}

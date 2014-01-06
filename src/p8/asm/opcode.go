package asm

import (
	"strconv"

	. "p8/opcode"
)

var opStr = map[uint16]string{
	Add:  "add",
	Addi: "addi",
	Sub:  "sub",
	Lui:  "lui",
	And:  "and",
	Andi: "andi",
	Or:   "or",
	Ori:  "ori",
	Xor:  "xor",
	Nor:  "nor",
	Slt:  "slt",
	Slti: "slti",
	Srl:  "srl",
	Sra:  "sra",
	Sllv: "sllv",
	Srlv: "srlv",
	Srav: "srav",

	Jr:   "jr",
	Beq:  "beq",
	Bne:  "bne",
	Mul:  "mul",
	Mulu: "mulu",
	Div:  "div",
	Divu: "divu",

	Ld:  "ld",
	Lw:  "lw",
	Lwu: "lwu",
	Lh:  "lu",
	Lhu: "lhu",
	Lb:  "lb",
	Lbu: "lbu",
	Sd:  "sd",
	Sw:  "sw",
	Sh:  "sh",
	Sb:  "sb",

	J:   "j",
	Jal: "jai",
}

func OpString(inst uint64) string {
	op := uint16(inst >> 48)

	if op&J != 0 {
		if op&Jal != 0 {
			return opStr[Jal]
		}
		return opStr[J]
	}

	s := opStr[op]
	if s == "" {
		return "nop" + strconv.Itoa(int(op))
	}

	return s
}

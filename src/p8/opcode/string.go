package opcode

import (
	"fmt"
)

func InstStr(i uint64) string {
	op, x, y, p, q, im := OpDec(i)

	if (op & J) != 0 {
		if (op & Jal) != 0 {
			return fmt.Sprintf("jal %016x", i<<3)
		}
		return fmt.Sprintf("j %016x", i<<3)
	}

	imu := uint64(im)
	ims := int64(int32(im))

	xpq := func(s string) string {
		return fmt.Sprintf("%s $%d, $%d, $%d", s, x, p, q)
	}
	xyiu := func(s string) string {
		return fmt.Sprintf("%s $%d, $%d, %d", s, x, y, imu)
	}
	xyis := func(s string) string {
		return fmt.Sprintf("%s $%d, $%d, %d", s, x, y, ims)
	}
	xypq := func(s string) string {
		return fmt.Sprintf("%s $%d, $%d, $%d, $%d", s, x, y, p, q)
	}

	switch op {
	case Halt:
		return "halt"
	case Rdtsc:
		return "rdtsc"
	case Rdttl:
		return "rdttl"

	case Add:
		return xpq("add")
	case Addi:
		return xyis("addi")
	case Sub:
		return xpq("sub")
	case Lui:
		return fmt.Sprintf("lui $%d, %d", x, imu)
	case And:
		return xpq("and")
	case Or:
		return xpq("or")
	case Ori:
		return xyiu("ori")
	case Xor:
		return xpq("xor")
	case Nor:
		return xpq("nor")
	case Slt:
		return xpq("slt")
	case Slti:
		return xyiu("slti")
	case Slli:
		return xyiu("slli")
	case Srli:
		return xyiu("srli")
	case Srai:
		return xyiu("srai")
	case Sll:
		return xpq("sll")
	case Srl:
		return xpq("srl")
	case Sra:
		return xpq("sra")

	case Jr:
		return fmt.Sprintf("jr $%d", p)

	case Beq:
		return fmt.Sprintf("beq %d", ims<<3)
	case Bne:
		return fmt.Sprintf("bne %d", ims<<3)

	case Mul:
		return xpq("mul")
	case Mulu:
		return xpq("mulu")
	case Div:
		return xypq("div")
	case Divu:
		return xypq("divu")

	case Ld:
		return xyis("ld")
	case Lw:
		return xyis("lw")
	case Lwu:
		return xyis("lwu")
	case Lh:
		return xyis("lh")
	case Lhu:
		return xyis("lhu")
	case Lb:
		return xyis("lb")
	case Lbu:
		return xyis("lbu")
	case Sd:
		return xyis("sd")
	case Sw:
		return xyis("sw")
	case Sh:
		return xyis("sh")
	case Sb:
		return xyis("sb")
	}

	return "noop"
}

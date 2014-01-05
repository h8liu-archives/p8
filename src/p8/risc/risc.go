package risc

// reg instructions
const (
	Halt = 0
	Jr   = iota
	Add
	Sub
	And
	Or
	Xor
	Nor
	Slt
	Sll
	Srl
	Sra
	Sllv
	Srlv
	Srav

	Mul = 0x10 + iota
	Mulu
	Div
	Divu
)

// immediate instructions
const (
	Addi = 0x10 + iota
	Andi
	Ori
	Slti
	Lw
	Lh
	Lhu
	Lb
	Lbu
	Lui
	Sw
	Sh
	Sb
	Beq
	Bne
)

// jump bits
const (
	J   = 0x80
	Jal = 0x40
)

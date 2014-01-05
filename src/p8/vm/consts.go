package vm

const (
	ExpNone = iota
	ExpHalt
	ExpAddr
	ExpPC
)

// reg instructions
const (
	halt = 0
	jr   = iota
	add
	sub
	and
	or
	xor
	nor
	slt
	sll
	srl
	sra
	sllv
	srlv
	srav

	mul = 0x10 + iota
	mulu
	div
	divu
)

// immediate instructions
const (
	addi = 0x10 + iota
	andi
	ori
	slti
	lw
	lh
	lhu
	lb
	lbu
	lui
	sw
	sh
	sb
	beq
	bne
)

const (
	_j   = 0x80
	_jal = 0x40
)

package asm

import (
	. "p8/opcode"
)

func (f *Func) Halt() *Inst                     { return f.I(Op(Halt)) }
func (f *Func) Rdtsc(x uint8) *Inst             { return f.I(OpX(Rdtsc, x)) }
func (f *Func) Rdttl(x uint8) *Inst             { return f.I(OpX(Rdttl, x)) }
func (f *Func) Add(x, p, q uint8) *Inst         { return f.I(OpXPQ(Add, x, p, q)) }
func (f *Func) Addi(x, y uint8, i int32) *Inst  { return f.I(OpXYIs(Addi, x, y, i)) }
func (f *Func) Sub(x, p, q uint8) *Inst         { return f.I(OpXPQ(Sub, x, p, q)) }
func (f *Func) Lui(x uint8, i uint32) *Inst     { return f.I(OpXI(Lui, x, i)) }
func (f *Func) And(x, p, q uint8) *Inst         { return f.I(OpXPQ(And, x, p, q)) }
func (f *Func) Andi(x, y uint8, i uint32) *Inst { return f.I(OpXYI(Andi, x, y, i)) }
func (f *Func) Or(x, p, q uint8) *Inst          { return f.I(OpXPQ(Or, x, p, q)) }
func (f *Func) Ori(x, y uint8, i uint32) *Inst  { return f.I(OpXYI(Ori, x, y, i)) }
func (f *Func) Xor(x, p, q uint8) *Inst         { return f.I(OpXPQ(Xor, x, p, q)) }
func (f *Func) Nor(x, p, q uint8) *Inst         { return f.I(OpXPQ(Xor, x, p, q)) }
func (f *Func) Slt(x, p, q uint8) *Inst         { return f.I(OpXPQ(Slt, x, p, q)) }
func (f *Func) Slti(x, y uint8, i uint32) *Inst { return f.I(OpXYI(Slti, x, y, i)) }
func (f *Func) Slli(x, y uint8, i uint32) *Inst { return f.I(OpXYI(Slli, x, y, i)) }
func (f *Func) Srli(x, y uint8, i uint32) *Inst { return f.I(OpXYI(Srli, x, y, i)) }
func (f *Func) Srai(x, y uint8, i uint32) *Inst { return f.I(OpXYI(Srai, x, y, i)) }
func (f *Func) Sll(x, p, q uint8) *Inst         { return f.I(OpXPQ(Sll, x, p, q)) }
func (f *Func) Srl(x, p, q uint8) *Inst         { return f.I(OpXPQ(Srl, x, p, q)) }
func (f *Func) Sra(x, p, q uint8) *Inst         { return f.I(OpXPQ(Sra, x, p, q)) }
func (f *Func) Jr(p uint8) *Inst                { return f.I(OpP(Jr, p)) }
func (f *Func) Beq(x, y uint8, l string) *Inst  { return f.I(OpXYI(Beq, x, y, 0)).L(l) }
func (f *Func) Bne(x, y uint8, l string) *Inst  { return f.I(OpXYI(Bne, x, y, 0)).L(l) }
func (f *Func) Mul(x, p, q uint8) *Inst         { return f.I(OpXPQ(Mul, x, p, q)) }
func (f *Func) Mulu(x, p, q uint8) *Inst        { return f.I(OpXPQ(Mulu, x, p, q)) }
func (f *Func) Div(x, y, p, q uint8) *Inst      { return f.I(OpXYPQ(Div, x, y, p, q)) }
func (f *Func) Divu(x, y, p, q uint8) *Inst     { return f.I(OpXYPQ(Divu, x, y, p, q)) }
func (f *Func) Ld(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Ld, x, y, i)) }
func (f *Func) Lw(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Lw, x, y, i)) }
func (f *Func) Lwu(x, y uint8, i int32) *Inst   { return f.I(OpXYIs(Lwu, x, y, i)) }
func (f *Func) Lh(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Lh, x, y, i)) }
func (f *Func) Lhu(x, y uint8, i int32) *Inst   { return f.I(OpXYIs(Lhu, x, y, i)) }
func (f *Func) Lb(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Lb, x, y, i)) }
func (f *Func) Sd(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Sd, x, y, i)) }
func (f *Func) Sw(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Sw, x, y, i)) }
func (f *Func) Sh(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Sh, x, y, i)) }
func (f *Func) Sb(x, y uint8, i int32) *Inst    { return f.I(OpXYIs(Sb, x, y, i)) }
func (f *Func) J(l string) *Inst                { return f.I(OpJ(J, 0)).L(l) }
func (f *Func) Jal(l string) *Inst              { return f.I(OpJ(J|Jal, 0)).L(l) }

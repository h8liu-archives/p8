package vm

import (
	. "p8/risc"
)

func (c *C) inst(i uint32) {
	op := uint8(i >> 24)
	r := c.gprs
	if (op & J) != 0 {
		c.pc = i << 2
		if (op & Jal) != 0 {
			r[15] = c.pc
		}
	} else {
		switch op >> 4 {
		case 0x0:
			c.regInst(i)
		case 0x1:
			c.immedInst(i)
		}
	}
}

func (c *C) regInst(i uint32) {
	r := c.gprs
	op := i >> 16
	x := (i >> 12) & 0xf
	y := (i >> 8) & 0xf
	p := (i >> 4) & 0xf
	q := i & 0xf

	switch op {
	case Halt:
		c.exp = ExpHalt
	case Jr:
		c.pc = r[p]
	case Add:
		r[x] = r[p] + r[q]
	case Sub:
		r[x] = r[p] - r[q]
	case And:
		r[x] = r[p] & r[q]
	case Or:
		r[x] = r[p] | r[q]
	case Xor:
		r[x] = r[p] ^ r[q]
	case Nor:
		r[x] = ^(r[p] | r[q])
	case Slt:
		r[x] = _slt(r[p], r[q])
	case Sll:
		r[x] = r[p] << q
	case Srl:
		r[x] = r[p] >> q
	case Sra:
		r[x] = uint32(int32(r[p]) >> q)
	case Sllv:
		r[x] = r[p] << r[q]
	case Srlv:
		r[x] = r[p] >> r[q]
	case Srav:
		r[x] = uint32(int32(r[p]) >> r[q])
	case Mul:
		t := int64(r[p]) * int64(r[q])
		r[x] = uint32(int32(t >> 32))
		r[y] = uint32(t)
	case Mulu:
		t := uint64(r[p]) * uint64(r[q])
		r[x] = uint32(t >> 32)
		r[y] = uint32(t)
	case Div:
		if r[q] == 0 {
			r[x], r[y] = 0, 0
		} else {
			r[x] = r[p] / r[q]
			r[x] = r[p] % r[q]
		}
	case Divu:
		_p := int32(r[p])
		_q := int32(r[q])
		if _q == 0 {
			r[x], r[y] = 0, 0
		} else {
			r[x] = uint32(_p / _q)
			r[y] = uint32(_p % _q)
		}
	}
}

func (c *C) immedInst(i uint32) {
	r := c.gprs

	op := i >> 24
	x := (i >> 20) & 0xf
	y := (i >> 16) & 0xf
	imu := i & 0xffff
	ims := se(uint16(imu))
	ad := r[y] + uint32(ims)

	switch op {
	case Addi:
		r[x] = r[y] + imu
	case Andi:
		r[x] = r[y] & imu
	case Ori:
		r[x] = r[y] | imu
	case Slti:
		r[x] = _slt(r[y], imu)
	case Lw:
		r[x] = c.rdw(ad)
	case Lh:
		r[x] = uint32(c.rdh(ad))
	case Lhu:
		r[x] = uint32(se(c.rdh(ad)))
	case Lb:
		r[x] = uint32(seb(c.rdb(ad)))
	case Lbu:
		r[x] = uint32(c.rdb(ad))
	case Lui:
		r[x] = (imu << 16) + (r[x] & 0xffff)
	case Sw:
		c.wrw(ad, r[x])
	case Sh:
		c.wrh(ad, uint16(r[x]))
	case Sb:
		c.wrb(ad, uint8(r[x]))
	case Beq:
		if r[x] == r[y] {
			c.pc += uint32(ims) << 2
		}
	case Bne:
		if r[x] != r[y] {
			c.pc += uint32(ims) << 2
		}
	}
}

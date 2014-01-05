package vm

import (
	. "p8/risc"
)

func (c *C) inst(i uint64) {
	op := i >> 48
	r := c.gprs
	if (op & J) != 0 {
		c.pc = i << 3
		if (op & Jal) != 0 {
			r[15] = c.pc
		}
	} else {
		switch op >> 12 {
		case 0:
			c.inst0(i)
		}
	}
}

func (c *C) inst0(i uint64) {
	r := c.gprs

	op := i >> 48
	x := (i >> 44) & 0xf
	y := (i >> 40) & 0xf
	p := (i >> 36) & 0xf
	q := (i >> 32) & 0xf

	imu := i & 0xffffffff
	ims := sew(uint32(imu))
	ad := r[y] + uint64(ims)

	switch op {
	// system
	case Halt:
		c.exp = ExpHalt
	case Rdtsc:
		r[x] = c.tsc

	// calculation
	case Add:
		r[x] = r[p] + r[q]
	case Addi:
		r[x] = r[y] + imu
	case Sub:
		r[x] = r[p] - r[q]
	case Lui:
		r[x] = (imu << 32) + (r[x] & 0xffffffff)
	case And:
		r[x] = r[p] & r[q]
	case Andi:
		r[x] = r[y] & imu
	case Or:
		r[x] = r[p] | r[q]
	case Ori:
		r[x] = r[y] | imu
	case Xor:
		r[x] = r[p] ^ r[q]
	case Nor:
		r[x] = ^(r[p] | r[q])
	case Slt:
		r[x] = _slt(r[p], r[q])
	case Slti:
		r[x] = _slt(r[y], imu)
	case Sll:
		r[x] = r[p] << q
	case Srl:
		r[x] = r[p] >> q
	case Sra:
		r[x] = uint64(int64(r[p]) >> q)
	case Sllv:
		r[x] = r[p] << r[q]
	case Srlv:
		r[x] = r[p] >> r[q]
	case Srav:
		r[x] = uint64(int64(r[p]) >> r[q])

	// control flow
	case Jr:
		c.pc = r[p]
	case Beq:
		if r[x] == r[y] {
			c.pc += uint64(ims) << 2
		}
	case Bne:
		if r[x] != r[y] {
			c.pc += uint64(ims) << 2
		}

	// mul and div
	case Mul:
		r[x] = uint64(int64(r[p]) * int64(r[q]))
	case Mulu:
		r[x] = r[p] * r[q]
	case Div:
		if r[q] == 0 {
			r[x], r[y] = 0, 0
		} else {
			r[x] = r[p] / r[q]
			r[x] = r[p] % r[q]
		}
	case Divu:
		_p := int64(r[p])
		_q := int64(r[q])
		if _q == 0 {
			r[x], r[y] = 0, 0
		} else {
			r[x] = uint64(_p / _q)
			r[y] = uint64(_p % _q)
		}

	// memory
	case Ld:
		r[x] = c.rdd(ad)
	case Lw:
		r[x] = uint64(sew(c.rdw(ad)))
	case Lwu:
		r[x] = uint64(c.rdh(ad))
	case Lh:
		r[x] = uint64(seh(c.rdh(ad)))
	case Lhu:
		r[x] = uint64(c.rdh(ad))
	case Lb:
		r[x] = uint64(seb(c.rdb(ad)))
	case Lbu:
		r[x] = uint64(c.rdb(ad))
	case Sd:
		c.wrd(ad, r[x])
	case Sw:
		c.wrw(ad, uint32(r[x]))
	case Sh:
		c.wrh(ad, uint16(r[x]))
	case Sb:
		c.wrb(ad, uint8(r[x]))
	}
}

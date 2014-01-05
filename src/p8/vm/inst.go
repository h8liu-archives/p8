package vm

func (c *C) inst(i uint32) {
	op := uint8(i >> 24)
	r := c.gprs
	if (op & _j) != 0 {
		c.pc = i << 2
		if (op & _jal) != 0 {
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
	case halt:
		c.exp = ExpHalt
	case jr:
		c.pc = r[p]
	case add:
		r[x] = r[p] + r[q]
	case sub:
		r[x] = r[p] - r[q]
	case and:
		r[x] = r[p] & r[q]
	case or:
		r[x] = r[p] | r[q]
	case xor:
		r[x] = r[p] ^ r[q]
	case nor:
		r[x] = ^(r[p] | r[q])
	case slt:
		r[x] = _slt(r[p], r[q])
	case sll:
		r[x] = r[p] << q
	case srl:
		r[x] = r[p] >> q
	case sra:
		r[x] = uint32(int32(r[p]) >> q)
	case sllv:
		r[x] = r[p] << r[q]
	case srlv:
		r[x] = r[p] >> r[q]
	case srav:
		r[x] = uint32(int32(r[p]) >> r[q])
	case mul:
		t := int64(r[p]) * int64(r[q])
		r[x] = uint32(int32(t >> 32))
		r[y] = uint32(t)
	case mulu:
		t := uint64(r[p]) * uint64(r[q])
		r[x] = uint32(t >> 32)
		r[y] = uint32(t)
	case div:
		if r[q] == 0 {
			r[x], r[y] = 0, 0
		} else {
			r[x] = r[p] / r[q]
			r[x] = r[p] % r[q]
		}
	case divu:
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
	case addi:
		r[x] = r[y] + imu
	case andi:
		r[x] = r[y] & imu
	case ori:
		r[x] = r[y] | imu
	case slti:
		r[x] = _slt(r[y], imu)
	case lw:
		r[x] = c.rdw(ad)
	case lh:
		r[x] = uint32(c.rdh(ad))
	case lhu:
		r[x] = uint32(se(c.rdh(ad)))
	case lb:
		r[x] = uint32(seb(c.rdb(ad)))
	case lbu:
		r[x] = uint32(c.rdb(ad))
	case lui:
		r[x] = (imu << 16) + (r[x] & 0xffff)
	case sw:
		c.wrw(ad, r[x])
	case sh:
		c.wrh(ad, uint16(r[x]))
	case sb:
		c.wrb(ad, uint8(r[x]))
	case beq:
		if r[x] == r[y] {
			c.pc += uint32(ims) << 2
		}
	case bne:
		if r[x] != r[y] {
			c.pc += uint32(ims) << 2
		}
	}
}

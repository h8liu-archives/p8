package vm

func u16(b []byte) uint16 {
	if b == nil {
		return 0
	}
	return uint16(b[0]) | uint16(b[1])<<8
}

func u16p(b []byte, v uint16) {
	if b == nil {
		return
	}
	b[0] = byte(v)
	b[1] = byte(v >> 8)
}

func u32(b []byte) uint32 {
	if b == nil {
		return 0
	}
	ret := uint32(b[0])
	ret |= uint32(b[1]) << 8
	ret |= uint32(b[2]) << 16
	ret |= uint32(b[3]) << 24
	return ret
}

func u32p(b []byte, v uint32) {
	if b == nil {
		return
	}
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func u64(b []byte) uint64 {
	if b == nil {
		return 0
	}
	ret := uint64(b[0])
	ret |= uint64(b[1]) << 8
	ret |= uint64(b[2]) << 16
	ret |= uint64(b[3]) << 24
	ret |= uint64(b[4]) << 32
	ret |= uint64(b[5]) << 40
	ret |= uint64(b[6]) << 48
	ret |= uint64(b[7]) << 56
	return ret
}

func u64p(b []byte, v uint64) {
	if b == nil {
		return
	}
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

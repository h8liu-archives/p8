package vm

func bin(c bool) uint32 {
	if c {
		return 1
	}
	return 0
}

func _slt(a, b uint32) uint32 {
	return bin(int32(a) < int32(b))
}

func se(c uint16) int32  { return int32(int16(c)) }
func ze(c uint16) uint32 { return uint32(c) }
func seb(c uint8) int32  { return int32(int8(c)) }
func zeb(c uint8) uint32 { return uint32(c) }

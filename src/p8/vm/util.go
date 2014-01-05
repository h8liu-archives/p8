package vm

func bin(c bool) uint64 {
	if c {
		return 1
	}
	return 0
}

func _slt(a, b uint64) uint64 {
	return bin(int64(a) < int64(b))
}

func sew(c uint32) int64  { return int64(int32(c)) }
func zew(c uint32) uint64 { return uint64(c) }
func seh(c uint16) int64  { return int64(int16(c)) }
func zeh(c uint16) uint64 { return uint64(c) }
func seb(c uint8) int64   { return int64(int8(c)) }
func zeb(c uint8) uint64  { return uint64(c) }

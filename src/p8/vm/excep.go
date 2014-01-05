package vm

const (
	ExcepHalt = 1 << iota
	ExcepAddr
	ExcepDeath
	ExcepMemPerm
)

package vm

// the virtual machine
type VM struct {
	r  []uint64
	pc uint64

	exp int
	tsc uint64

	mem []byte
}

func New(memSize int) *VM {
	if memSize%1024 != 0 {
		panic("memory not aligned to 1K")
	}
	if memSize == 0 {
		panic("zero memory")
	}

	ret := new(VM)
	ret.r = make([]uint64, 16)
	ret.mem = make([]byte, memSize)

	return ret
}

func (vm *VM) step() int {
	// read
	i := vm.rdd(vm.pc)
	vm.pc += 8

	// exec
	vm.inst(i)

	// clean up
	vm.r[0] = 0
	u64p(vm.mem[0:8], 0)
	vm.tsc++

	return vm.exp
}

func (vm *VM) Restart(start uint64) int {
	vm.ClearTSC()
	return vm.ResumeAt(start)
}

func (vm *VM) ResumeAt(start uint64) int {
	vm.pc = start
	return vm.Resume()
}

func (vm *VM) Resume() int {
	vm.exp = ExcepNone
	for vm.exp == ExcepNone {
		vm.step()
	}

	return vm.exp
}

func (vm *VM) Step() int {
	vm.exp = ExcepNone
	return vm.step()
}

func (vm *VM) StepAt(start uint64) int {
	vm.pc = start
	return vm.Step()
}

func (vm *VM) Load(m []byte, offset uint64) {
	if offset%8 != 0 {
		panic("offset not aligned")
	}
	n := uint64(len(m))
	copy(vm.mem[offset:offset+n], m[:n])
}

func (vm *VM) ClearTSC() { vm.tsc = 0 }

func (vm *VM) PC() uint64       { return vm.pc }
func (vm *VM) TSC() uint64      { return vm.tsc }
func (vm *VM) R(a uint8) uint64 { return vm.r[a] }

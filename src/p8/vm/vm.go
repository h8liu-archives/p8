package vm

// the virtual machine
type VM struct {
	r []uint64

	pc  uint64
	TSC uint64
	TTL uint64

	mem []byte
	e   uint64
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

func (vm *VM) except(e uint64) {
	vm.e |= e
}

func (vm *VM) tick() {
	vm.TSC++
	if vm.TTL > 0 {
		vm.TTL--
		if vm.TTL == 0 {
			vm.except(ExcepDeath)
		}
	}
}

func (vm *VM) step() uint64 {
	// read
	i := vm.rdd(vm.pc)
	vm.pc += 8

	// exec
	vm.inst(i)

	// clean up
	vm.r[0] = 0
	u64p(vm.mem[0:8], 0)

	vm.tick()

	return vm.e
}

func (vm *VM) ResumeAt(start uint64) uint64 {
	vm.pc = start
	return vm.Resume()
}

func (vm *VM) Resume() uint64 {
	vm.e = 0
	for vm.e == 0 {
		vm.step()
	}

	return vm.e
}

func (vm *VM) Step() uint64 {
	vm.e = 0
	return vm.step()
}

func (vm *VM) StepAt(start uint64) uint64 {
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

func (vm *VM) R(a uint8) uint64 { return vm.r[a] }
func (vm *VM) PC() uint64       { return vm.pc }

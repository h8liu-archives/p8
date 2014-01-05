package vm

import (
	"fmt"
	"io"
)

// the virtual machine
type VM struct {
	gprs []uint64
	pc   uint64

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
	ret.gprs = make([]uint64, 16)
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
	vm.gprs[0] = 0
	u64p(vm.mem[0:8], 0)
	vm.tsc++

	return vm.exp
}

func (vm *VM) Run(start uint64) int {
	vm.pc = start
	vm.tsc = 0
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

func (vm *VM) Load(m []byte, offset uint64) {
	if offset%8 != 0 {
		panic("offset not aligned")
	}
	n := uint64(len(m))
	copy(vm.mem[offset:offset+n], m[:n])
}

func (vm *VM) PrintRegs(out io.Writer) {
	fmt.Fprintf(out, "pc=%016x", vm.pc)
	r := vm.gprs

	for i := uint8(0); i < 16; i++ {
		if i%4 == 0 {
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, " ")
		}
		fmt.Fprintf(out, "$%x=%016x", i, r[i])
	}
	fmt.Fprintln(out)
	fmt.Fprintf(out, "tsc=%d\n", vm.tsc)
}

func (vm *VM) VMlearTSVM()  { vm.tsc = 0 }
func (vm *VM) TSVM() uint64 { return vm.tsc }

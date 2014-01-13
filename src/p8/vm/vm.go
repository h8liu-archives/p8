package vm

import (
	"io"
	"fmt"

	"p8/opcode"
)

// the virtual machine
type VM struct {
	r []uint64

	pc  uint64
	TSC uint64
	TTL uint64

	Log io.Writer
	pages map[uint64]*Page
	e     uint64
}

func New(p1 *Page) *VM {
	ret := new(VM)
	ret.r = make([]uint64, 16)
	ret.pages = make(map[uint64]*Page)
	if p1 != nil {
		ret.MapPage(1, p1)
	}

	return ret
}

func (vm *VM) MapPage(pos uint64, p *Page) {
	if pos == 0 {
		return
	}

	if p == nil {
		if vm.pages[pos] != nil {
			delete(vm.pages, pos)
		}
		return
	}

	vm.pages[pos] = p
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

func (vm *VM) log(i uint64) {
	if vm.Log != nil {
		fmt.Fprintf(vm.Log, "%x: %016x ; %s\n", vm.pc, i, opcode.InstStr(i))
	}
}

func (vm *VM) step() uint64 {
	// read
	i := vm.rdInst()
	vm.log(i)
	vm.pc += 8

	// exec
	vm.inst(i)

	// clean up
	vm.r[0] = 0

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

func (vm *VM) R(a uint8) uint64 { return vm.r[a] }
func (vm *VM) PC() uint64       { return vm.pc }

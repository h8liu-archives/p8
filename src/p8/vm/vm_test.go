package vm

import (
	"bytes"
	. "p8/bin"
	. "p8/opcode"
	"testing"
)

func TestBasicOps(t *testing.T) {
	p := func(insts []uint64) *Page {
		ret := new(bytes.Buffer)
		buf := make([]byte, 8)
		for _, inst := range insts {
			U64p(buf, inst)
			ret.Write(buf)
		}

		buf = ret.Bytes()
		n := len(buf)
		if n > PageSize {
			panic("too large")
		}

		page := NewPage(PermAll)
		copy(page.Bytes[:n], buf)

		return page
	}

	v := func(insts ...uint64) *VM {
		ret := New(p(insts))
		ret.TTL = 100
		ret.ResumeAt(PageStart(1))
		return ret
	}

	reg := func(vm *VM, r uint8, v uint64) {
		got := vm.R(r)
		if got != v {
			t.Errorf("r%d expect %016x, got %016x", r, v, got)
		}
	}

	vm := v(OpXYI(Addi, 1, 0, 30))
	reg(vm, 0, 0)
	reg(vm, 1, 30)

	vm = v(
		OpXYI(Addi, 1, 0, 3),
		OpXYI(Addi, 2, 0, 4),
		OpXPQ(Add, 3, 1, 2),
	)
	reg(vm, 3, 7)

	vm = v(
		OpXYI(Addi, 5, 0, 13),
		OpXYI(Addi, 12, 0, 5),
		OpXPQ(Sub, 5, 5, 12),
	)
	reg(vm, 5, 8)

	vm = v(
		OpJ(Jal, 4096+8),
	)
	reg(vm, 15, 4096+8)
}

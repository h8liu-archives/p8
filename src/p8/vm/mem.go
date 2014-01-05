package vm

func (vm *VM) ad(a uint64, n uint, perm uint64) []byte {
	ad := (a >> n) << n

	if ad != a {
		vm.except(ExcepAddr)
		return nil
	}

	p := vm.pages[a>>PageOffset]
	if p == nil {
		vm.except(ExcepAddr)
		return nil
	}

	if !p.HavePerm(perm) {
		vm.except(ExcepMemPerm)
		return nil
	}

	return p.S(a, 1<<n)
}

func (vm *VM) rdInst() uint64 { return u64(vm.ad(vm.pc, 3, PermExec)) }

func (vm *VM) rdd(a uint64) uint64 { return u64(vm.ad(a, 3, 0)) }
func (vm *VM) rdw(a uint64) uint32 { return u32(vm.ad(a, 2, 0)) }
func (vm *VM) rdh(a uint64) uint16 { return u16(vm.ad(a, 1, 0)) }
func (vm *VM) rdb(a uint64) uint8  { return vm.ad(a, 0, 0)[0] }

func (vm *VM) wrd(a uint64, v uint64) { u64p(vm.ad(a, 3, PermWrite), v) }
func (vm *VM) wrw(a uint64, v uint32) { u32p(vm.ad(a, 2, PermWrite), v) }
func (vm *VM) wrh(a uint64, v uint16) { u16p(vm.ad(a, 1, PermWrite), v) }
func (vm *VM) wrb(a uint64, v uint8)  { vm.ad(a, 0, PermWrite)[0] = v }

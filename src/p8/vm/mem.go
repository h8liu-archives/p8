package vm

func (vm *VM) ad(a uint64, n uint) uint64 {
	// TODO: page handling
	ad := (a >> n) << n
	if ad != a {
		vm.exp = ExcepAddr
	}
	return ad
}

func (vm *VM) _rdd(a uint64) uint64 { return u64(vm.mem[a : a+8]) }
func (vm *VM) _rdw(a uint64) uint32 { return u32(vm.mem[a : a+4]) }
func (vm *VM) _rdh(a uint64) uint16 { return u16(vm.mem[a : a+2]) }
func (vm *VM) _rdb(a uint64) uint8  { return vm.mem[a] }

func (vm *VM) _wrd(a uint64, v uint64) { u64p(vm.mem[a:a+8], v) }
func (vm *VM) _wrw(a uint64, v uint32) { u32p(vm.mem[a:a+4], v) }
func (vm *VM) _wrh(a uint64, v uint16) { u16p(vm.mem[a:a+4], v) }
func (vm *VM) _wrb(a uint64, v uint8)  { vm.mem[a] = v }

func (vm *VM) rdd(a uint64) uint64 { return vm._rdd(vm.ad(a, 3)) }
func (vm *VM) rdw(a uint64) uint32 { return vm._rdw(vm.ad(a, 2)) }
func (vm *VM) rdh(a uint64) uint16 { return vm._rdh(vm.ad(a, 1)) }
func (vm *VM) rdb(a uint64) uint8  { return vm._rdb(vm.ad(a, 0)) }

func (vm *VM) wrd(a uint64, v uint64) { vm._wrd(vm.ad(a, 3), v) }
func (vm *VM) wrw(a uint64, v uint32) { vm._wrw(vm.ad(a, 2), v) }
func (vm *VM) wrh(a uint64, v uint16) { vm._wrh(vm.ad(a, 1), v) }
func (vm *VM) wrb(a uint64, v uint8)  { vm._wrb(vm.ad(a, 0), v) }

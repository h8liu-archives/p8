package vm

func (c *C) ad(a uint64, n uint) uint64 {
	// TODO: page handling
	ad := (a >> n) << n
	if ad != a {
		c.exp = ExpAddr
	}
	return ad
}

func (c *C) _rdd(a uint64) uint64 { return u64(c.mem[a : a+8]) }
func (c *C) _rdw(a uint64) uint32 { return u32(c.mem[a : a+4]) }
func (c *C) _rdh(a uint64) uint16 { return u16(c.mem[a : a+2]) }
func (c *C) _rdb(a uint64) uint8  { return c.mem[a] }

func (c *C) _wrd(a uint64, v uint64) { u64p(c.mem[a:a+8], v) }
func (c *C) _wrw(a uint64, v uint32) { u32p(c.mem[a:a+4], v) }
func (c *C) _wrh(a uint64, v uint16) { u16p(c.mem[a:a+4], v) }
func (c *C) _wrb(a uint64, v uint8)  { c.mem[a] = v }

func (c *C) rdd(a uint64) uint64 { return c._rdd(c.ad(a, 3)) }
func (c *C) rdw(a uint64) uint32 { return c._rdw(c.ad(a, 2)) }
func (c *C) rdh(a uint64) uint16 { return c._rdh(c.ad(a, 1)) }
func (c *C) rdb(a uint64) uint8  { return c._rdb(c.ad(a, 0)) }

func (c *C) wrd(a uint64, v uint64) { c._wrd(c.ad(a, 3), v) }
func (c *C) wrw(a uint64, v uint32) { c._wrw(c.ad(a, 2), v) }
func (c *C) wrh(a uint64, v uint16) { c._wrh(c.ad(a, 1), v) }
func (c *C) wrb(a uint64, v uint8)  { c._wrb(c.ad(a, 0), v) }

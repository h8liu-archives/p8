package vm

import (
	"encoding/binary"
)

var enc = binary.LittleEndian

func (c *C) memlen() uint32 { return uint32(len(c.mem)) }

func (c *C) adw(a uint32) uint32 {
	ad := (a >> 2) << 2
	if ad != a {
		c.exp = ExpAddr
	}
	if ad+4 > c.memlen() {
		c.exp = ExpAddr
		return 0
	}
	return ad
}
func (c *C) adh(a uint32) uint32 {
	ad := (a >> 1) << 1
	if ad != a {
		c.exp = ExpAddr
	}
	if ad+2 > c.memlen() {
		c.exp = ExpAddr
		return 0
	}
	return ad
}

func (c *C) adb(a uint32) uint32 {
	if a >= c.memlen() {
		c.exp = ExpAddr
		return 0
	}
	return a
}

func (c *C) _rdw(a uint32) uint32 { return enc.Uint32(c.mem[a : a+4]) }
func (c *C) _rdh(a uint32) uint16 { return enc.Uint16(c.mem[a : a+2]) }
func (c *C) _rdb(a uint32) uint8  { return c.mem[a] }

func (c *C) _wrw(a uint32, v uint32) { enc.PutUint32(c.mem[a:a+4], v) }
func (c *C) _wrh(a uint32, v uint16) { enc.PutUint16(c.mem[a:a+4], v) }
func (c *C) _wrb(a uint32, v uint8)  { c.mem[a] = v }

func (c *C) rdw(a uint32) uint32 { return c._rdw(c.adw(a)) }
func (c *C) rdh(a uint32) uint16 { return c._rdh(c.adh(a)) }
func (c *C) rdb(a uint32) uint8  { return c._rdb(c.adb(a)) }

func (c *C) wrw(a uint32, v uint32) { c._wrw(c.adw(a), v) }
func (c *C) wrh(a uint32, v uint16) { c._wrh(c.adh(a), v) }
func (c *C) wrb(a uint32, v uint8)  { c._wrb(c.adb(a), v) }

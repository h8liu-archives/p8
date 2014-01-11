package asm

type Label struct {
	Name string
	Pos  uint64
}

func newLabel(n string) *Label {
	if len(n) == 0 {
		panic("bug")
	}

	return &Label{Name: n}
}

func (self *Label) P(p uint64) *Label {
	self.Pos = p
	return self
}

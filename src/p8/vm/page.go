package vm

const (
	PageOffset  = 12
	PageSize    = 1 << PageOffset
	PageOffMask = PageSize - 1
)

const (
	PermWrite = 1 << iota
	PermExec

	PermAll = PermWrite | PermExec
)

type Page struct {
	perm uint64
	b    []byte
}

func PageStart(p uint64) uint64 { return p << PageOffset }

func NewPage(perm uint64) *Page {
	ret := new(Page)
	ret.perm = perm
	ret.b = make([]byte, PageSize)
	return ret
}

func (p *Page) S(o uint64, n uint) []byte {
	o &= PageOffMask
	return p.b[o : o+uint64(n)]
}

func (p *Page) HavePerm(perm uint64) bool {
	return p.perm&perm == perm
}

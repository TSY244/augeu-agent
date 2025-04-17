package engUtils

type Base struct {
}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) SizeForStr(str []string) int {
	return len(str)
}

package engUtils

type Base struct {
}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) SizeForStr(str []string) int {
	return len(str)
}

func (b *Base) GeneFileSegmentation(chunkSize int, src string) string {
	var dst string
	for i := 0; i < chunkSize; i++ {
		dst += src
	}
	return dst
}

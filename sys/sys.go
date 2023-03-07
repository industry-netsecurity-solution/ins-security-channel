package sys

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func Memset(dst []byte, offset, size int, v byte) int {
	if dst == nil {
		return 0
	}
	var i int
	count := min(len(dst), size)
	for i = 0; i < count; i++ {
		dst[i] = v
	}

	return i
}

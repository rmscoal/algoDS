package arrays

func rotLeft(a []int32, d int32) []int32 {
	return append(a[d:], a[:d]...)
}

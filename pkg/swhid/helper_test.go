package swhid

func repeat(b byte, n int) []byte {
	result := []byte{}
	for i := 0; i < n; i++ {
		result = append(result, b)
	}
	return result
}

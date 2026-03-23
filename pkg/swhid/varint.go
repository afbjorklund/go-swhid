//go:build sql

package swhid

/*
** from sqlite3:compressFunc (and uncompressFunc)
**
** The output is a BLOB that begins with a variable-length integer that
** is the input size in bytes (the size of X before compression).  The
** variable-length integer is implemented as 1 to 5 bytes.  There are
** seven bits per integer stored in the lower seven bits of each byte.
** More significant bits occur first.  The most significant bit (0x80)
** is a flag to indicate the end of the integer.
 */

func varint(n int) []byte {
	var i, j int
	x := make([]uint, 8)
	for i = 4; i >= 0; i-- {
		x[i] = uint((n >> (7 * (4 - i))) & 0x7f)
	}
	for i = 0; i < 4 && x[i] == 0; i++ {
	}
	p := make([]byte, 5)
	for j = 0; i <= 4; i++ {
		p[j] = byte(x[i])
		j++
	}
	p[j-1] |= 0x80
	return p[:j]
}

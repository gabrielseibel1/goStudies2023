package lc7

import (
	"strconv"
)

func reverse(x int) int {
	s := strconv.Itoa(x)
	r := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	if r[len(r)-1] == byte('-') {
		r = r[:len(r)-1]
		r = append([]byte{'-'}, r...)
	}
	y, err := strconv.ParseInt(string(r), 10, 32)
	if err != nil {
		return 0
	}
	return int(y)
}

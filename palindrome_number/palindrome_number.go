package palindromenumber

import (
	"math"
	"strconv"
	"strings"
)

func isPalindromeUsingStrings(x int) bool {
	s := strconv.Itoa(x)
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	r, _ := strconv.Atoi(b.String())
	return (x == r)
}

func IsPalindrome(x int) bool {
	d := int(math.Log10(float64(x)))
	var r int
	for p, q := 0, d; p <= d && q >= 0; p, q = p+1, q-1 {
		r += int(math.Pow10(q)) * (x / int(math.Pow10(p)))
	}
	return x == r
}

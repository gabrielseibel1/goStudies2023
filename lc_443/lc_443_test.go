package lc443

import "testing"

func Test_compress(t *testing.T) {
	t.Run("aabbccc", func(t *testing.T) {
		exp := "a2b2c3"
		testString(t, []byte("aabbccc"), exp)
	})

	t.Run("a", func(t *testing.T) {
		exp := "a"
		testString(t, []byte("a"), exp)
	})

	t.Run("abbbbbbbbbbbb", func(t *testing.T) {
		exp := "ab12"
		testString(t, []byte("abbbbbbbbbbbb"), exp)
	})
}

func testString(t *testing.T, chars []byte, exp string) {
	size := compress(chars)
	copyChars := make([]byte, size)
	copySize := copy(copyChars, chars)
	if size != len(exp) || copySize != size {
		t.Error("size")
	}
	if string(copyChars) != exp {
		t.Errorf("string (%s)", copyChars)
	}
}

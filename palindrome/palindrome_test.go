package palindrome

import (
	"testing"
)

func testLongestPalindromeString(s string, want string, t *testing.T) {
	res := LongestPalindrome(s)
	if res != want {
		t.Fatalf(`LongestPalindrome(%q) = %q, want match for %#q, nil`, s, res, want)
	}
}

func TestLongestPalindromeA(t *testing.T) {
	testLongestPalindromeString("a", "a", t)
}

func TestLongestPalindromeBabad(t *testing.T) {
	testLongestPalindromeString("babad", "aba", t)
}

func TestLongestPalindromeAaaaa(t *testing.T) {
	testLongestPalindromeString("aaaaa", "aaaaa", t)
}

func TestLongestPalindromeCbbd(t *testing.T) {
	testLongestPalindromeString("cbbd", "bb", t)
}

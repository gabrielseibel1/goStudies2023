package palindrome

func longestPalindromeMemo(s string, memo map[string]string) string {
	if v, e := memo[s]; e {
		return v
	}

	if s == "" {
		memo[s] = s
		return memo[s]
	}

	endsEqual := s[0] == s[len(s)-1]
	if len(s) <= 3 && endsEqual {
		memo[s] = s
		return memo[s]
	}

	left, mid, right := s[:len(s)-1], s[1:len(s)-1], s[1:]

	palMid := longestPalindromeMemo(mid, memo)
	palLeft := longestPalindromeMemo(left, memo)
	palRight := longestPalindromeMemo(right, memo)

	if palMid == mid && endsEqual {
		memo[s] = s
	} else if len(palMid) >= len(palLeft) && len(palMid) >= len(palRight) {
		memo[s] = palMid
	} else if len(palLeft) >= len(palMid) && len(palLeft) >= len(palRight) {
		memo[s] = palLeft
	} else if len(palRight) >= len(palMid) && len(palRight) >= len(palLeft) {
		memo[s] = palRight
	}

	return memo[s]
}

func longestPalindromeTabu(s string, tab [][]bool) string {
	longest := ""

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] && (i == j || (i+1 < len(s) && j-1 >= 0 && tab[i+1][j-1])) {
				tab[i][j] = true
				if j-i >= len(longest) {
					longest = s[i : j+1]
				}
			}
		}
	}
	return longest
}

func LongestPalindrome(s string) string {
	// return longestPalindromeMemo(s, make(map[string]string))
	// init and seed values
	tab := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		tab[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			if i == j || j == i-1 {
				tab[i][j] = true
			} else {
				tab[i][j] = false
			}
		}
	}
	return longestPalindromeTabu(s, tab)
}

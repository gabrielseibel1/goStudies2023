package lc17

var letters = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations := make([]string, 0)
	combination := make([]byte, 0, len(digits))
	fillLetterCombinations(digits, 0, &combination, &combinations)
	return combinations
}

func fillLetterCombinations(digits string, start int, combination *[]byte, combinations *[]string) {
	if len(*combination) == len(digits) {
		*combinations = append(*combinations, string(*combination))
		return
	}
	for i := start; i < len(digits); i++ {
		for _, l := range letters[digits[i]] {
			*combination = append(*combination, l)
			fillLetterCombinations(digits, i+1, combination, combinations)
			*combination = (*combination)[:len(*combination)-1]
		}
	}
}

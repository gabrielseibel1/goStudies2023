package aws

import (
	"sort"
)

// func main() {
// 	// println(findReviewScore("GoodProductbutScrapAfterWash", []string{"crap", "odpro"}))
// 	println(findReviewScore("FastDeliveryOkayProduct", []string{"eryoka", "yo", "eli"}))
// 	// for _, p := range permutations([]string{"eryoka", "yo", "eli"}) {
// 	// 	println(strings.Join(p, ","))
// 	// }
// }

func FindReviewScore(review string, prohibitedWords []string) int32 {
	var result int32

	i, j := 0, 0
	for j < len(review) {
		maxJ := j
		for _, pw := range prohibitedWords {
			startPW := j - len(pw)
			if review[j-len(pw):j+1] == pw {
				i = startPW + 1
				j = i
			} else {
				maxJ++
				if l := maxJ - i; l > int(result) {
					result = int32(l)
				}
			}
		}
	}

	return result
}

/*
 * Complete the 'findReviewScore' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING review
 *  2. STRING_ARRAY prohibitedWords
 */

// type strRange struct {
// 	i int
// 	j int
// }
// func findReviewScore(review string, prohibitedWords []string) int32 {
// 	var result int32

// 	reviewLen := len(review)

// 	sort.Slice(prohibitedWords, func(i, j int) bool {
// 		return len(prohibitedWords[i]) < len(prohibitedWords[j])
// 	})

// 	for i, j := 0, 0; i < reviewLen && j < reviewLen; {
// 		var maxCanAdvanceJ int
// 		var foundWord string
// 		for _, pw := range prohibitedWords {
// 			for k := j; k-j < len(pw) && k < reviewLen; {
// 				k++
// 				if review[j:k] != pw[0:k-j] {
// 					if l := len(pw); l > maxCanAdvanceJ {
// 						if j+l > len(review) {
// 							maxCanAdvanceJ = reviewLen - j
// 						} else {
// 							maxCanAdvanceJ = l
// 						}
// 					}
// 					break
// 				}
// 				if k-j == len(pw) {
// 					foundWord = pw
// 					maxCanAdvanceJ = len(pw) - 1
// 				}
// 			}
// 		}
// 		if l := j + maxCanAdvanceJ - i; l > int(result) {
// 			result = int32(l)
// 		}
// 		if foundWord != "" {
// 			i = j + 1
// 			j++
// 		} else {
// 			j += maxCanAdvanceJ
// 		}
// 	}

// 	return result
// }

// func findReviewScore(review string, prohibitedWords []string) int32 {
// 	if review == "" {
// 		return 0
// 	}
// 	var m int32
// 	review = strings.ToLower(review)
// 	for _, prohibitedWordsPerm := range permutations(prohibitedWords) {
// 		substrings := map[string]struct{}{review: {}}
// 		for _, p := range prohibitedWordsPerm {
// 			p = strings.ToLower(p)
// 			newSubstrings := map[string]struct{}{}
// 			for substring := range substrings {
// 				substring = strings.Replace(substring, p, "*", -1)
// 				for _, newSubstring := range strings.Split(substring, "*") {
// 					newSubstrings[newSubstring] = struct{}{}
// 				}
// 			}
// 			substrings = newSubstrings
// 		}
// 		for k := range substrings {
// 			if l := int32(len(k)); l > m {
// 				m = l
// 			}
// 		}
// 	}
// 	return m
// }

// func permutations(ss []string) [][]string {
// 	p := make([][]string, 0)
// 	for i := range ss {
// 		tt := make([]string, len(ss))
// 		_ = copy(tt, ss)
// 		for j := range ss {
// 			tt[i] = tt[j]
// 			tt[j] = ss[i]
// 			p = append(p, innerPermutations(ss, j)...)
// 		}
// 	}
// 	return p
// }

// func innerPermutations(ss []string, fixed int) [][]string {
// 	l := permutations(ss[:fixed])
// 	r := permutations(ss[fixed+1:])
// 	lm := append(l, []string{ss[fixed]})
// 	lmr := append(lm, r...)
// 	return lmr
// }

/*
 * Complete the 'getMaxPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY frontend
 *  2. INTEGER_ARRAY backend
 */
func getMaxPairs(frontend []int32, backend []int32) int32 {
	var p int32
	sort.Slice(frontend, func(i, j int) bool {
		return frontend[i] < frontend[j]
	})
	sort.Slice(backend, func(i, j int) bool {
		return backend[i] < backend[j]
	})
	j := 0
	for i := range frontend {
		if frontend[i] > backend[j] {
			p++
			j++
		}
	}
	return p
}

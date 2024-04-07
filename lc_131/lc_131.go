package lc131

import (
	"slices"
	"strings"
)

func partition(s string) [][]string {
	partitions := make([][]string, 0)
	partition := make([]string, 0)
	fillPartitions(s, &partitions, &partition, 0)
	return partitions
}

func fillPartitions(s string, partitions *[][]string, partition *[]string, i int) {

	if isCompleteWord(*partition, s) {
		*partitions = append(*partitions, slices.Clone(*partition))
		return
	}

	if i >= len(s) {
		return
	}

	// push letters here
	letters := make([]byte, 0)

	// when find palindrome, try to complete partition from there
	for j := i; j < len(s); j++ {
		letters = append(letters, s[j])
		if isPalindrome(letters) {
			// forward
			*partition = append(*partition, string(letters))

			// recurse
			fillPartitions(s, partitions, partition, j+1)

			// backtrack
			*partition = (*partition)[:len(*partition)-1]
		}
	}
}

func isPalindrome(letters []byte) bool {
	for i, j := 0, len(letters)-1; i <= j; i, j = i+1, j-1 {
		if letters[i] != letters[j] {
			return false
		}
	}
	return true
}

func isCompleteWord(partition []string, s string) bool {
	for _, p := range partition {
		if strings.HasPrefix(s, p) {
			s = s[len(p):]
			if len(s) == 0 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

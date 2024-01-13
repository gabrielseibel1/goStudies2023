package lc_347_top_k_frequent_elements

import (
	"github.com/gabrielseibel1/fungo/conv"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	c := make(map[int]int)
	for _, e := range nums {
		c[e] += 1
	}
	o := conv.MapToPairs(c)
	sort.Slice(o, func(i, j int) bool {
		return o[i].V > o[j].V
	})
	s := conv.PairsKeysToSlice(o)
	return s[:k]
}

// time	: O(n+k*log(k))
// space: O(k)

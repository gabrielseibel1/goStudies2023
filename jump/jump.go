package jump

// https://leetcode.com/problems/jump-game-ii/

func jumpTabu(nums []int, tab []int) int {
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j <= i+nums[i] && j < len(tab); j++ {
			jumps := tab[i] + 1
			if jumps < tab[j] || tab[j] == 0 {
				tab[j] = jumps
			}
		}
	}
	return tab[len(tab)-1]
}

func Jump(nums []int) int {
	return jumpTabu(nums, make([]int, len(nums)))
}

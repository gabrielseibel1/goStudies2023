package lc268

func missingNumber(nums []int) int {
	ref := 0
	for i := 0; i <= len(nums); i++ {
		ref += i
	}
	act := 0
	for _, v := range nums {
		act += v
	}
	return ref - act
}

package lc287

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}
}

// func findDuplicate(nums []int) int {
// 	sp := uint64(0)

// 	for _, n := range nums {
// 		if a := (sp >> (n - 1)) % 2; a == 1 {
// 			return n
// 		} else {
// 			b := uint64(1 << (n - 1))
// 			sp += b
// 		}
// 	}
// 	panic("this should never happen")
// }

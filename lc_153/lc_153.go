package lc153

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	i := len(nums) / 2

	if nums[l] <= nums[r] {
		return nums[l]
	}

	size := len(nums) / 2
	for {
		// corner case of reaching extremities
		if i == l || i == r {
			return nums[i]
		}
		// 4, 5, (6), 1, 2, 3
		if i+1 <= r && nums[i] > nums[i+1] {
			return nums[i+1]
		}
		// 4, 5, 6, (1), 2, 3
		if i-1 >= l && nums[i] < nums[i-1] {
			return nums[i]
		}
		// is left of edge
		if nums[i] > nums[r] {
			// go slightly right
			size /= 2
			if size == 0 {
				size++
			}
			i += size
			continue
		}
		// is right of edge
		if nums[i] < nums[l] {
			// go slightly left
			size /= 2
			if size == 0 {
				size++
			}
			i -= size
			continue
		}
	}
}

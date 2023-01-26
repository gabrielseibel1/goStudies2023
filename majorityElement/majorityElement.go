package majorityelement

// https://leetcode.com/problems/majority-element

func MajorityElement(nums []int) int {
	return majority(nums, 0, len(nums)-1)
}

func majority(nums []int, from, to int) int {
	if from == to {
		return nums[from]
	}

	mean := (from + to) / 2
	majLeft := majority(nums, from, mean)
	majRight := majority(nums, mean+1, to)

	if majLeft == majRight {
		return majLeft
	}

	countMajLeft, countMajRight := count(majLeft, majRight, nums, from, to)
	if countMajLeft >= countMajRight {
		return majLeft
	}
	return majRight
}

func count(a, b int, nums []int, from, to int) (int, int) {
	countA, countB := 0, 0
	for i := from; i <= to; i++ {
		if nums[i] == a {
			countA = countA + 1
		}
		if nums[i] == b {
			countB = countB + 1
		}
	}
	return countA, countB
}

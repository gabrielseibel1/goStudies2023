package containermostwater

func MaxArea(height []int) int {
	i, j := 0, len(height)-1
	maxContained, prevIH, prevJH := 0, 0, 0
	for i < j {
		maxContained = max(maxContained, min(height[i], height[j])*(j-i))

		prevIH, prevJH = height[i], height[j]

		if prevIH <= prevJH {
			for !(height[i] > prevIH) && i+1 <= j {
				i++
			}
		} else {
			for !(height[j] > prevJH) && j-1 >= i {
				j--
			}
		}
	}
	return maxContained
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

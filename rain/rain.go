package rain

// https://leetcode.com/problems/trapping-rain-water/

func Trap(height []int) int {
	water := 0

	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	for i := 0; i < len(maxLeft); i++ {
		if i == 0 {
			maxLeft[i] = 0
			continue
		}
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}

	for i := len(maxRight) - 1; i >= 0; i-- {
		if i == len(maxRight)-1 {
			maxRight[i] = 0
			continue
		}
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	for i := 0; i < len(height); i++ {
		water += max(min(maxLeft[i], maxRight[i])-height[i], 0)
	}
	return water
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// func Trap(height []int) int {
// 	parentheses := make([]byte, 0, len(height))
// 	levels := make([]int, 0, len(height))
// 	positions := make([]int, 0, len(height))

// 	for i := 0; i < len(height)-1; i++ {
// 		dif := height[i+1] - height[i]
// 		nPar := int(math.Abs(float64(dif)))
// 		var c byte
// 		var level int
// 		if dif < 0 { // downhill
// 			c = '('
// 			level = max(height[i+1], height[i]) - 1
// 		} else if dif > 0 { // uphill
// 			c = ')'
// 			level = min(height[i+1], height[i])
// 		}
// 		for j := 0; j < nPar; j++ {
// 			parentheses = append(parentheses, c)
// 			positions = append(positions, i)
// 			levels = append(levels, level)
// 			if dif < 0 { // downhill
// 				level -= 1
// 			} else if dif > 0 { // uphill
// 				level += 1
// 			}
// 		}
// 	}

// 	distance := 0
// 	seenLevels := make(map[int]bool)
// 	for _, level := range levels {
// 		if seen, exists := seenLevels[level]; exists && seen {
// 			continue
// 		}
// 		stackPar := make([]byte, 0, len(parentheses))
// 		stackPos := make([]int, 0, len(parentheses))
// 		for i, c := range parentheses {
// 			if levels[i] == level {
// 				stackPar = append(stackPar, c)
// 				stackPos = append(stackPos, positions[i])
// 				if len(stackPar) >= 2 && stackPar[len(stackPar)-2] == '(' && stackPar[len(stackPar)-1] == ')' {
// 					distance += int(math.Abs(float64(stackPos[len(stackPos)-1] - stackPos[len(stackPos)-2])))
// 					stackPar = stackPar[:len(stackPar)-2]
// 					stackPos = stackPos[:len(stackPos)-2]
// 				}
// 			}
// 		}
// 		seenLevels[level] = true
// 	}
// 	return distance
// }

package lc443

import "strconv"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	size := 0
	current := 0
	for current < len(chars) {
		i := current
		for {
			l := i - current
			if i == len(chars) || chars[i] != chars[current] {

				chars[size] = chars[current]
				size++
				if l > 1 {
					lStr := strconv.Itoa(l)
					for j := 0; j < len(lStr); j++ {
						chars[size] = lStr[j]
						size++
					}
				}

				current = i
				break
			}
			i++
		}
	}
	return size
}

package main

import (
	"strings"
)

func MatrixSpiral(strArr []string) string {
	n := len(strArr)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return at(strArr, 0, 0)
	}
	result := make([]string, 0, n*n)
	for layer := 0; layer <= n/2; layer++ {
		row, col := layer, layer
		for i := col; i < n-layer; i++ {
			col = i
			result = append(result, at(strArr, row, col))
		}
		for i := row + 1; i < n-layer; i++ {
			row = i
			result = append(result, at(strArr, row, col))
		}
		for i := col - 1; i >= layer; i-- {
			col = i
			result = append(result, at(strArr, row, col))
		}
		for i := row - 1; i > layer; i-- {
			row = i
			result = append(result, at(strArr, row, col))
		}
	}
	return strings.Join(result, ",")
}

func at(strArray []string, row, col int) string {
	str := strArray[row]
	str = str[1 : len(str)-1]
	nums := strings.Split(str, ",")
	return nums[col]
}

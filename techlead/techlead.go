package techlead

import (
	"fmt"
	"slices"
)

// input word = "banana"
// longest pal= "anana"

// funcao de validar palindromo
// incremental de forma que tendo um palindromo eu possa validar se uma superstring dela eh palindromo tb

// aaabbbccc
// abcacba -> a

// mapa de freqs -> vazio -> vazio, uma letra #1 -> a letra,

// aaaabbbbcccc
// aabbccccbbaa

// mapa de frequencias
// se eu tiver um numero par de contagem de uma letra
// posso espelhar todos os caracteres dessa letra
// se eu tiver contagem impar eu posso usar o que sobrar no meio

func LongestPalindrome(s string) string {
	m := make(map[byte]int)
	p := make([]byte, 0)
	for i := range s {
		m[s[i]] = m[s[i]] + 1
	}

	var other string
	for k := range m {
		if m[k]%2 == 0 { // eh par
			for i := 0; i < m[k]/2; i++ {
				// coloca no palindromo
				p = append([]byte{k}, append(p, k)...)
			}
		} else {
			for i := 0; i < m[k]/2; i++ {
				// coloca no palindromo
				p = append([]byte{k}, append(p, k)...)
			}
			// reserva a letra para preencher meio do palindromo
			other = string([]byte{k})
			fmt.Println("other is", other)
		}
	}
	if other != "" {
		p = append(p[:len(p)/2], append([]byte(other), p[(len(p)/2):]...)...)
	}

	return string(p)
}

type position struct {
	r int
	c int
}

type path []position

func ShortestPathFromAtoB(grid [][]int) path {
	/*
		0 1 0 0
		0 1 0 0
		0 0 0 0
		1 1 0 B

		A 1 1 1 0
		0 0 0 1 0
		0 1 0 0 0
		0 1 1 1 0
		0 0 0 0 B
	*/

	end := position{len(grid) - 1, len(grid) - 1}

	// BFS eu tenho, para cada elemento da fronteira (n), seu path (n * n)
	// DFS eu tenho, para cada elemento da fronteira (1),seu path (n)

	q := make([]path, 1)
	q[0] = path{{0, 0}}
	size := len(q)
	for size > 0 {
		for i := 0; i < size; i++ {
			pth := q[0]
			tail := pth[len(pth)-1]
			grid[tail.r][tail.c] = 1
			q = q[1:] // grab from the start and put others in the end
			// enqueue neighbors of the current
			for _, neigh := range neighborsOfTail(grid, tail) {
				fmt.Println("neighs of", tail, neigh)
				newPath := append(slices.Clone(pth), neigh)
				if neigh == end {
					return newPath
				}
				q = append(q, newPath)
			}
		}
		size = len(q)
	}
	panic("no paths found from A to B")
}

func neighborsOfTail(grid [][]int, tail position) []position {
	neighbors := make([]position, 0)
	l := len(grid)
	n := position{tail.r + 1, tail.c}
	if inGrid(l, n) && grid[n.r][n.c] == 0 {
		neighbors = append(neighbors, n)
	}
	n = position{tail.r - 1, tail.c}
	if inGrid(l, n) && grid[n.r][n.c] == 0 {
		neighbors = append(neighbors, n)
	}
	n = position{tail.r, tail.c + 1}
	if inGrid(l, n) && grid[n.r][n.c] == 0 {
		neighbors = append(neighbors, n)
	}
	n = position{tail.r, tail.c - 1}
	if inGrid(l, n) && grid[n.r][n.c] == 0 {
		neighbors = append(neighbors, n)
	}
	return neighbors
}

func inGrid(l int, p position) bool {
	return p.r >= 0 && p.r < l && p.c >= 0 && p.c < l
}

// poderia representar o path diferente ???

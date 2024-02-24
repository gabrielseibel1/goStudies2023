package lc_49_group_anagrams

import "sort"

func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]int)
	for i, s := range strs {
		c := code(s)
		m[c] = append(m[c], i)
	}
	r := make([][]string, 0)
	for _, v := range m {
		s := make([]string, 0)
		for _, i := range v {
			s = append(s, strs[i])
		}
		r = append(r, s)
	}
	return r
}

func code(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// -- faster than sorting: counting letters:

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[letters][]int)
//	for i, s := range strs {
//		l := lettersIn(s)
//		m[l] = append(m[l], i)
//	}
//	r := make([][]string, 0)
//	for _, v := range m {
//		s := make([]string, 0)
//		for _, i := range v {
//			s = append(s, strs[i])
//		}
//		r = append(r, s)
//	}
//	return r
//}
//
//func lettersIn(s string) letters {
//	l := letters{}
//	for _, c := range s {
//		switch c {
//		case 'a':
//			l.a += 1
//		case 'b':
//			l.b += 1
//		case 'c':
//			l.c += 1
//		case 'd':
//			l.d += 1
//		case 'e':
//			l.e += 1
//		case 'f':
//			l.f += 1
//		case 'g':
//			l.g += 1
//		case 'h':
//			l.h += 1
//		case 'i':
//			l.i += 1
//		case 'j':
//			l.j += 1
//		case 'k':
//			l.k += 1
//		case 'l':
//			l.l += 1
//		case 'm':
//			l.m += 1
//		case 'n':
//			l.n += 1
//		case 'o':
//			l.o += 1
//		case 'p':
//			l.p += 1
//		case 'q':
//			l.q += 1
//		case 'r':
//			l.r += 1
//		case 's':
//			l.s += 1
//		case 't':
//			l.t += 1
//		case 'u':
//			l.u += 1
//		case 'v':
//			l.v += 1
//		case 'w':
//			l.w += 1
//		case 'x':
//			l.x += 1
//		case 'y':
//			l.y += 1
//		case 'z':
//			l.z += 1
//		}
//	}
//	return l
//}
//
//type letters struct {
//	a int
//	b int
//	c int
//	d int
//	e int
//	f int
//	g int
//	h int
//	i int
//	j int
//	k int
//	l int
//	m int
//	n int
//	o int
//	p int
//	q int
//	r int
//	s int
//	t int
//	u int
//	v int
//	w int
//	x int
//	y int
//	z int
//}

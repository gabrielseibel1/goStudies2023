package main

type Set interface {
	Add(entry int64)
	Entries() []int64
}

func main() {
	var set Set = &MySet{}
	go set.Add(1)
	go set.Add(1)
	go set.Add(2)
	go set.Add(3)
	go set.Add(5)

	go set.Entries()
}

type MySet struct {
	seen    map[int64]struct{}
	entries []int64
}

func (s *MySet) Add(entry int64) {
	if s.seen == nil {
		s.seen = map[int64]struct{}{}
	}
	if _, ok := s.seen[entry]; !ok {
		s.seen[entry] = struct{}{}
		s.entries = append(s.entries, entry)
	}
}

func (s *MySet) Entries() []int64 {
	return s.entries
}

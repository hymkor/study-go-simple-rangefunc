package main

type seq struct {
	values []int
}

func newSeq(v ...int) *seq {
	return &seq{values: v}
}

func (s *seq) Each(callback func(int) bool) {
	for _, v := range s.values {
		if !callback(v) {
			break
		}
	}
}

func main() {
	s := newSeq(1, 3, 5, 7, 9)
	s.Each(func(v int) bool {
		println(v)
		return true
	})
}

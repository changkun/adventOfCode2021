package main

import (
	"sort"

	"changkun.de/x/adventOfCode2021/utils"
)

type stack struct {
	b []byte
}

func newStack() *stack {
	return &stack{b: []byte{}}
}

func (s *stack) push(b byte) {
	s.b = append(s.b, b)
}

func (s *stack) pop() byte {
	if len(s.b) == 0 {
		return 0
	}

	b := s.b[len(s.b)-1]
	s.b = s.b[:len(s.b)-1]
	return b
}

func (s *stack) len() int {
	return len(s.b)
}

func isRight(b byte) bool {
	if b == ')' || b == ']' || b == '}' || b == '>' {
		return true
	}
	return false
}

func isPair(a, b byte) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	if a == '<' && b == '>' {
		return true
	}
	return false
}

func read(fname string) (int, int) {
	r, _ := utils.Read(fname)

	illegals := map[byte]int{
		')': 0,
		']': 0,
		'}': 0,
		'>': 0,
	}
	all := []int{}

xx:
	for r.Scan() {
		tt := r.Text()
		s := newStack()
		p2score := 0

		for i := 0; i < len(tt); i++ {
			if !isRight(tt[i]) {
				s.push(tt[i])
				continue
			}

			b := s.pop()
			if isPair(b, tt[i]) {
				continue
			}

			illegals[tt[i]]++
			continue xx
		}

		for {
			b := s.pop()
			if b == 0 {
				break
			}
			if b == '(' {
				p2score = p2score*5 + 1
			}
			if b == '[' {
				p2score = p2score*5 + 2
			}
			if b == '{' {
				p2score = p2score*5 + 3
			}
			if b == '<' {
				p2score = p2score*5 + 4
			}
		}

		all = append(all, p2score)
	}

	sum := illegals[')']*3 + illegals[']']*57 + illegals['}']*1197 + illegals['>']*25137
	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})
	return sum, all[(len(all)-1)/2]
}

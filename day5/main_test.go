package main_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

type vec2 struct {
	x, y int
}

func newVec2(s string) vec2 {
	vs := strings.Split(s, ",")
	x, _ := strconv.Atoi(vs[0])
	y, _ := strconv.Atoi(vs[1])
	return vec2{x, y}
}

type mat struct {
	n     int
	elems []int
}

func newMat(n int) mat {
	return mat{n: n, elems: make([]int, n*n)}
}

func (m *mat) add(i, j, v int) {
	m.elems[i*m.n+j] += v
}

func (m *mat) ranging(v1, v2 vec2, diagnal bool) {
	if v1.x != v2.x && v1.y != v2.y {
		if !diagnal {
			return
		}

		if v1.x < v2.x && v1.y < v2.y {
			// v1 is bottom left
			i := v1.x
			j := v1.y
			for {
				if i > v2.x && j > v2.y {
					return
				}
				m.add(i, j, 1)
				i++
				j++
			}
		} else if v1.x < v2.x && v1.y > v2.y {
			// v1 is bottom top
			i := v1.x
			j := v1.y
			for {
				if i > v2.x && j < v2.y {
					return
				}
				m.add(i, j, 1)
				i++
				j--
			}
		} else if v1.x > v2.x && v1.y < v2.y {
			// v1 is bottom right
			i := v2.x
			j := v2.y
			for {
				if i > v1.x && j < v1.y {
					return
				}
				m.add(i, j, 1)
				i++
				j--
			}
		} else if v1.x > v2.x && v1.y > v2.y {
			// v1 is top right
			i := v2.x
			j := v2.y
			for {
				if i > v1.x && j > v1.y {
					return
				}
				m.add(i, j, 1)
				i++
				j++
			}
		}
		return
	}
	if v1.x == v2.x && v1.y == v2.y {
		m.add(v1.x, v2.x, 1)
		return
	}

	if v1.x == v2.x {
		for i := int(math.Min(float64(v1.y), float64(v2.y))); i <= int(math.Max(float64(v1.y), float64(v2.y))); i++ {
			m.add(v1.x, i, 1)
		}
		return
	}

	if v1.y == v2.y {
		for i := int(math.Min(float64(v1.x), float64(v2.x))); i <= int(math.Max(float64(v1.x), float64(v2.x))); i++ {
			m.add(i, v1.y, 1)
		}
		return
	}

	panic("impossible")
}

func (m *mat) counting() (n int) {
	for _, e := range m.elems {
		if e > 1 {
			n++
		}
	}
	return
}

func calc(t *testing.T, fname string, diagnal bool) int {
	s := utils.Read(t, fname)

	max := 0
	lines := [][2]vec2{}
	for s.Scan() {
		tt := s.Text()
		fields := strings.Fields(tt)
		v1 := newVec2(fields[0])
		v2 := newVec2(fields[2])
		if v1.x > max {
			max = v1.x
		}
		if v1.y > max {
			max = v1.y
		}
		if v2.x > max {
			max = v2.x
		}
		if v2.y > max {
			max = v2.y
		}
		lines = append(lines, [2]vec2{v1, v2})
	}

	m := newMat(max + 1)
	for _, l := range lines {
		v1, v2 := l[0], l[1]
		m.ranging(v1, v2, diagnal)
	}
	return m.counting()
}

func TestPart1(t *testing.T) {
	n := calc(t, "test.txt", false)
	if n != 5 {
		t.Fatalf("want 5 got: %v", n)
	}

	t.Log(calc(t, "input.txt", false))
}

func TestPart2(t *testing.T) {
	n := calc(t, "test.txt", true)
	if n != 12 {
		t.Fatalf("want 12 got: %v", n)
	}

	t.Log(calc(t, "input.txt", true))
}

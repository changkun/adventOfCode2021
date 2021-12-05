package main

import (
	"math"
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

type vec2 struct{ x, y int }

func newVec2(s string) vec2 {
	vs := strings.Split(s, ",")
	x, _ := strconv.Atoi(vs[0])
	y, _ := strconv.Atoi(vs[1])
	return vec2{x, y}
}
func (v *vec2) inc()            { v.x, v.y = v.x+1, v.y+1 }
func (v *vec2) dec()            { v.x, v.y = v.x+1, v.y-1 }
func (v vec2) meet(u vec2) bool { return v.x == u.x && v.y == u.y }
func (v vec2) sortX(u vec2) (vec2, vec2) {
	if v.x < u.x {
		return v, u
	}
	return u, v
}
func (v vec2) diagnal(u vec2) bool {
	dx := v.x - u.x
	dy := v.y - u.y
	return math.Abs(float64(dy)/float64(dx)) == 1
}
func max(ns ...int) int {
	m := math.MinInt
	for _, n := range ns {
		m = int(math.Max(float64(m), float64(n)))
	}
	return m
}
func min(ns ...int) int {
	m := math.MaxInt
	for _, n := range ns {
		m = int(math.Min(float64(m), float64(n)))
	}
	return m
}

type mat struct {
	n     int
	elems []int
}

func newMat(n int) mat {
	return mat{n: n, elems: make([]int, n*n)}
}

func (m *mat) add(i, j, v int) { m.elems[i*m.n+j] += v }
func (m *mat) drawLine(v1, v2 vec2, drawDiagnal bool) {
	switch {
	case v1.x == v2.x && v1.y == v2.y:
		m.add(v1.x, v1.y, 1)
		return
	case v1.x == v2.x:
		for i := min(v1.y, v2.y); i <= max(v1.y, v2.y); i++ {
			m.add(v1.x, i, 1)
		}
		return
	case v1.y == v2.y:
		for i := min(v1.x, v2.x); i <= max(v1.x, v2.x); i++ {
			m.add(i, v1.y, 1)
		}
		return
	default:
		if !drawDiagnal || !v1.diagnal(v2) {
			return
		}

		if (v1.x < v2.x && v1.y < v2.y) || (v1.x > v2.x && v1.y > v2.y) {
			v1, v2 = v1.sortX(v2)
			for v := v1; !v.meet(v2); v.inc() {
				m.add(v.x, v.y, 1)
			}
		} else if (v1.x < v2.x && v1.y > v2.y) || (v1.x > v2.x && v1.y < v2.y) {
			v1, v2 = v1.sortX(v2)
			for v := v1; !v.meet(v2); v.dec() {
				m.add(v.x, v.y, 1)
			}
		}
	}
}

func (m *mat) counting() (n int) {
	for _, e := range m.elems {
		if e > 1 {
			n++
		}
	}
	return
}

func calc(fname string, drawDiagnal bool) int {
	s, _ := utils.Read(fname)
	n, lines := 0, [][2]vec2{}
	for s.Scan() {
		tt := s.Text()
		fields := strings.Fields(tt)
		v1, v2 := newVec2(fields[0]), newVec2(fields[2])
		n = max(n, v1.x, v2.x, v1.y, v2.y)
		lines = append(lines, [2]vec2{v1, v2})
	}

	m := newMat(n + 1)
	for _, l := range lines {
		v1, v2 := l[0], l[1]
		m.drawLine(v1, v2, drawDiagnal)
	}
	return m.counting()
}

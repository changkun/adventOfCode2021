package main

import (
	"sort"
	"testing"
)

func TestPart1(t *testing.T) {
	m := ReadMatrix("input.txt")

	locals := 0
	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			if m.Local(i, j) {
				// t.Log(i, j, locals, m.elems[i][j])
				locals += m.elems[i][j] + 1
			}
		}
	}

	t.Log(locals)
}

type Pos struct {
	x, y int
}

func searchHorizonal(m Mat, p []Pos) []Pos {
	horizontal := []Pos{}

	for _, pp := range p {
		horizontal = append(horizontal, pp)
		u := pp.x
		v := pp.y
		for j := v; j >= 0; j-- {
			if m.elems[u][j] == 9 {
				break
			}
			horizontal = append(horizontal, Pos{u, j})
		}
		for j := v; j < m.n; j++ {
			if m.elems[u][j] == 9 {
				break
			}
			horizontal = append(horizontal, Pos{u, j})
		}
	}
	return horizontal
}

func searchVertical(m Mat, p []Pos) []Pos {
	vertical := []Pos{}
	for _, pp := range p {
		vertical = append(vertical, pp)

		x := pp.x
		y := pp.y
		for i := x; i >= 0; i-- {
			if m.elems[i][y] == 9 {
				break
			}
			vertical = append(vertical, Pos{i, y})
		}
		for i := x; i < m.m; i++ {
			if m.elems[i][y] == 9 {
				break
			}
			vertical = append(vertical, Pos{i, y})
		}
	}
	return vertical
}

func TestPart2(t *testing.T) {
	m := ReadMatrix("input.txt")

	locals := []Pos{}
	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			if m.Local(i, j) {
				locals = append(locals, Pos{i, j})
			}
		}
	}

	ret := []int{}
	for _, local := range locals {

		region := map[Pos]int{local: 1}
		lastN := len(region)
		for {
			ps := []Pos{}
			for p := range region {
				ps = append(ps, p)
			}

			horizontal := searchHorizonal(m, ps)
			for _, p := range horizontal {
				region[p]++
			}

			vertical := searchVertical(m, horizontal)
			for _, p := range vertical {
				region[p]++
			}

			currentN := len(region)
			if currentN == lastN {
				break
			}
			lastN = currentN
		}
		ret = append(ret, len(region))
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] > ret[j]
	})

	t.Log(ret[0] * ret[1] * ret[2])
}

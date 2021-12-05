package main

import (
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

type mat5 struct {
	skip  bool
	elems [25]int
}

func newMat5(ns ...int) mat5 {
	m := mat5{}
	for i := range m.elems {
		m.elems[i] = ns[i]
	}
	return m
}

// Draw draws a matched element to -1
func (m *mat5) Draw(elem int) {
	for i, n := range m.elems {
		if n == elem {
			m.elems[i] = -1
		}
	}
}

func (m *mat5) At(i, j int) int {
	return m.elems[5*i+j]
}

func (m *mat5) SumRow(row int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += m.At(row, i)
	}
	return sum
}
func (m *mat5) SumCol(col int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += m.At(i, col)
	}
	return sum
}

func (m *mat5) IsWin() bool {
	for i := 0; i < 5; i++ {
		if m.SumCol(i) == -5 || m.SumRow(i) == -5 {
			return true
		}
	}
	return false
}

func (m *mat5) SumPositive() int {
	sum := 0
	for i := 0; i < len(m.elems); i++ {
		if m.elems[i] < 0 {
			continue
		}
		sum += m.elems[i]
	}
	return sum
}

func prepareInputs(fname string) ([]int, []mat5) {
	s, _ := utils.Read(fname)

	// 1. prepare input numbers
	rawInput := s.Text()
	ns := strings.Split(rawInput, ",")
	drawers := make([]int, len(ns))
	for i := range drawers {
		n, _ := strconv.Atoi(ns[i])
		drawers[i] = n
	}

	// 2. prepare input matrices
	mats := []mat5{}
	cache := []int{}
	for s.Scan() {
		tt := s.Text()
		if tt == "" {
			continue
		}

		ns := strings.Fields(tt)
		for i := range ns {
			n, _ := strconv.Atoi(ns[i])
			cache = append(cache, n)
		}
	}
	for i := 0; i < len(cache); i += 25 {
		mats = append(mats, newMat5(cache[i:i+25]...))
	}
	return drawers, mats
}

func computeWin(fname string) (winMats []mat5, winDrawers []int) {
	// figure out every possible win orders.
	drawers, mats := prepareInputs(fname)
	winMats = []mat5{}
	winDrawers = []int{}
	for j := 0; j < len(drawers); j++ {
		for i := range mats {
			mats[i].Draw(drawers[j])
		}

		for i := range mats {
			if mats[i].IsWin() && !mats[i].skip {
				winMats = append(winMats, mats[i])
				winDrawers = append(winDrawers, drawers[j])
				mats[i].skip = true
			}
		}
	}
	return
}

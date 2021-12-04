package main_test

import (
	"strconv"
	"strings"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

type Mat5 struct {
	skip  bool
	elems [25]int
}

func NewMat5(ns ...int) Mat5 {
	m := Mat5{skip: false}
	for i := range m.elems {
		m.elems[i] = ns[i]
	}
	return m
}

// Draw draws a matched element to -1
func (m *Mat5) Draw(elem int) {
	for i, n := range m.elems {
		if n == elem {
			m.elems[i] = -1
		}
	}
}

func (m *Mat5) At(i, j int) int {
	return m.elems[5*i+j]
}

func (m *Mat5) SumRow(row int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += m.At(row, i)
	}
	return sum
}
func (m *Mat5) SumCol(col int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += m.At(i, col)
	}
	return sum
}

func (m *Mat5) IsWin() bool {
	for i := 0; i < 5; i++ {
		if m.SumCol(i) == -5 {
			return true
		}
		if m.SumRow(i) == -5 {
			return true
		}
	}
	return false
}

func (m *Mat5) Skip() bool {
	return m.skip
}

func (m *Mat5) SetSkip() {
	m.skip = true
}

func (m *Mat5) SumPositive() int {
	sum := 0
	for i := 0; i < len(m.elems); i++ {
		if m.elems[i] < 0 {
			continue
		}
		sum += m.elems[i]
	}
	return sum
}

func prepareInputs(t *testing.T, fname string) ([]int, []Mat5) {
	s := utils.Read(t, fname)
	if !s.Scan() {
		t.Fatal("unexpected input")
	}

	// 1. prepare input numbers
	rawInput := s.Text()
	ns := strings.Split(rawInput, ",")
	drawers := make([]int, len(ns))
	for i := range drawers {
		n, err := strconv.Atoi(ns[i])
		if err != nil {
			t.Fatal(err)
		}
		drawers[i] = n
	}

	// 2. prepare input matrices
	mats := []Mat5{}
	cache := []int{}
	for s.Scan() {
		tt := s.Text()
		if tt == "" {
			continue
		}

		ns := strings.Fields(tt)
		for i := range ns {
			n, err := strconv.Atoi(ns[i])
			if err != nil {
				t.Fatal(err)
			}
			cache = append(cache, n)
		}
	}
	for i := 0; i < len(cache); i += 25 {
		m := NewMat5(cache[i : i+25]...)
		mats = append(mats, m)
	}
	return drawers, mats
}

func computeWin(t *testing.T, fname string) (winMats []Mat5, winDrawers []int) {
	// figure out every possible win orders.
	drawers, mats := prepareInputs(t, fname)
	winMats = []Mat5{}
	winDrawers = []int{}
	for j := 0; j < len(drawers); j++ {
		for i := range mats {
			mats[i].Draw(drawers[j])
		}

		for i := range mats {
			if mats[i].IsWin() && !mats[i].Skip() {
				winMats = append(winMats, mats[i])
				winDrawers = append(winDrawers, drawers[j])
				mats[i].SetSkip()
			}
		}
	}
	return
}

func TestPart1(t *testing.T) {
	winMats, winDrawers := computeWin(t, "test.txt")
	v := winMats[0].SumPositive() * winDrawers[0]
	if v != 4512 {
		t.Fatalf("want 4512 got %v", v)
	}

	winMats, winDrawers = computeWin(t, "input.txt")
	t.Log(winMats[0].SumPositive() * winDrawers[0])
}

func TestPart2(t *testing.T) {
	winMats, winDrawers := computeWin(t, "test.txt")
	v := winMats[len(winMats)-1].SumPositive() * winDrawers[len(winDrawers)-1]
	if v != 1924 {
		t.Fatalf("want 1924 got %v", v)
	}

	winMats, winDrawers = computeWin(t, "input.txt")
	t.Log(winMats[len(winMats)-1].SumPositive() * winDrawers[len(winDrawers)-1])
}

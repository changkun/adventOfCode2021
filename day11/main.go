package main

import (
	"log"
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

type mat10 struct {
	elems [][]int
}

func newMat10(elems [][]int) mat10 {
	return mat10{elems: elems}
}

func (m mat10) String() string {
	b := strings.Builder{}
	for i := 0; i < len(m.elems); i++ {
		nn := m.elems[i]
		s := ""
		for j := 0; j < len(nn); j++ {
			s += strconv.Itoa(nn[j])
		}
		b.Write([]byte(s + "\n"))
	}
	return b.String()
}

func (m *mat10) inc() (flashed int, synced bool) {
	zeroedPos := [][2]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m.elems[i][j] = (m.elems[i][j] + 1) % 10
			if m.elems[i][j] == 0 {
				zeroedPos = append(zeroedPos, [2]int{i, j})
			}
		}
	}

again:
	flashed += len(zeroedPos)

	// sync?
	sum := 0
	for i := range m.elems {
		for j := range m.elems[i] {
			sum += m.elems[i][j]
		}
	}
	if sum == 0 {
		synced = true
		return
	}

	newZeroedPos := [][2]int{}
	for _, pos := range zeroedPos {
		i := pos[0]
		j := pos[1]

		if i-1 >= 0 {
			if j-1 >= 0 && m.elems[i-1][j-1] != 0 {
				m.elems[i-1][j-1] = (m.elems[i-1][j-1] + 1) % 10
				if m.elems[i-1][j-1] == 0 {
					newZeroedPos = append(newZeroedPos, [2]int{i - 1, j - 1})
				}
			}
			if m.elems[i-1][j] != 0 {
				m.elems[i-1][j] = (m.elems[i-1][j] + 1) % 10
				if m.elems[i-1][j] == 0 {
					newZeroedPos = append(newZeroedPos, [2]int{i - 1, j})
				}
			}
			if j+1 < 10 && m.elems[i-1][j+1] != 0 {
				m.elems[i-1][j+1] = (m.elems[i-1][j+1] + 1) % 10
				if m.elems[i-1][j+1] == 0 {
					newZeroedPos = append(newZeroedPos, [2]int{i - 1, j + 1})
				}
			}
		}
		if j-1 >= 0 && m.elems[i][j-1] != 0 {
			m.elems[i][j-1] = (m.elems[i][j-1] + 1) % 10
			if m.elems[i][j-1] == 0 {
				newZeroedPos = append(newZeroedPos, [2]int{i, j - 1})
			}
		}
		if j+1 < 10 && m.elems[i][j+1] != 0 {
			m.elems[i][j+1] = (m.elems[i][j+1] + 1) % 10
			if m.elems[i][j+1] == 0 {
				newZeroedPos = append(newZeroedPos, [2]int{i, j + 1})
			}
		}
		if i+1 < 10 {
			if j-1 >= 0 && m.elems[i+1][j-1] != 0 {
				m.elems[i+1][j-1] = (m.elems[i+1][j-1] + 1) % 10
				if m.elems[i+1][j-1] == 0 {
					newZeroedPos = append(newZeroedPos, [2]int{i + 1, j - 1})
				}
			}
			if m.elems[i+1][j] != 0 {
				m.elems[i+1][j] = (m.elems[i+1][j] + 1) % 10
				if m.elems[i+1][j] == 0 {
					newZeroedPos = append(newZeroedPos, [2]int{i + 1, j})
				}
			}
			if j+1 < 10 && m.elems[i+1][j+1] != 0 {
				m.elems[i+1][j+1] = (m.elems[i+1][j+1] + 1) % 10
				if m.elems[i+1][j+1] == 0 {
					newZeroedPos = append(newZeroedPos, [2]int{i + 1, j + 1})
				}
			}
		}
	}

	if len(newZeroedPos) != 0 {
		zeroedPos = newZeroedPos
		goto again
	}
	return
}

func part1(fname string) {
	s, _ := utils.Read(fname)
	col := [][]int{}
	for s.Scan() {
		tt := s.Text()
		ns := strings.Split(tt, "")
		row := []int{}
		for i := range ns {
			n, _ := strconv.Atoi(ns[i])
			row = append(row, n)
		}
		col = append(col, row)
	}
	m := newMat10(col)

	n := 0
	for i := 0; i < 100; i++ {
		flashes, _ := m.inc()
		n += flashes
	}
	log.Println(n)
}

func part2(fname string) {
	s, _ := utils.Read(fname)
	col := [][]int{}
	for s.Scan() {
		tt := s.Text()
		ns := strings.Split(tt, "")
		row := []int{}
		for i := range ns {
			n, _ := strconv.Atoi(ns[i])
			row = append(row, n)
		}
		col = append(col, row)
	}
	m := newMat10(col)

	i := 0
	for {
		i++
		_, synced := m.inc()
		if synced {
			break
		}
	}
	log.Println(i)
}

func main() {
	log.SetFlags(0)
	part1("test.txt")
	part1("input.txt")
	part2("test.txt")
	part2("input.txt")
}

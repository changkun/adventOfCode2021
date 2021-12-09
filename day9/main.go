package main

import (
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

type Mat struct {
	elems [][]int
	m, n  int
}

func (m *Mat) Local(i, j int) bool {
	cur := m.elems[i][j]

	// left
	if i-1 >= 0 && m.elems[i-1][j] <= cur {
		return false
	}

	// right
	if i+1 < m.m && m.elems[i+1][j] <= cur {
		return false
	}

	// top
	if j-1 >= 0 && m.elems[i][j-1] <= cur {
		return false
	}

	// bottom
	if j+1 < m.n && m.elems[i][j+1] <= cur {
		return false
	}

	return true
}

func ReadMatrix(fname string) Mat {
	s, _ := utils.Read(fname)

	ret := [][]int{}
	for s.Scan() {
		tt := s.Text()

		row := []int{}
		ns := strings.Split(tt, "")
		for i := range ns {
			n, _ := strconv.Atoi(ns[i])
			row = append(row, n)
		}
		ret = append(ret, row)
	}

	return Mat{elems: ret, m: len(ret), n: len(ret[0])}
}

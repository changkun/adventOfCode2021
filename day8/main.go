package main

import (
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

func uniqueCount(fname string) int {
	s, _ := utils.Read(fname)
	n := 0

	for s.Scan() {
		tt := s.Text()
		patterns := strings.Split(tt, " | ")
		digits := strings.Split(patterns[1], " ")

		for _, d := range digits {
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				n++
			}
		}
	}

	return n
}

func decode(known ...string) string {
	all := map[int][]string{
		1: {},
		7: {},
		4: {},
		8: {},
	}

	for _, k := range known {

		switch len(k) {
		case 2:
			Perm(k, func(s string) {
				all[1] = append(all[1], s)
			})
		case 3:
			Perm(k, func(s string) {
				all[7] = append(all[7], s)
			})
		case 4:
			Perm(k, func(s string) {
				all[4] = append(all[4], s)
			})
		case 7:
			Perm(k, func(s string) {
				all[8] = append(all[8], s)
			})
		}
	}

	allCombinations := [][]string{}

	for {
		stop := true
		for k := range all {
			if len(all[k]) != 0 {
				stop = false
				break
			}
		}
		if stop {
			break
		}

		comb := []string{}
		for k := range all {
			if len(all[k]) > 0 {
				comb = append(comb, all[k][len(all[k])-1])
				all[k] = all[k][:len(all[k])-1]
			}
		}
		allCombinations = append(allCombinations, comb)
	}

	for _, comb := range allCombinations {
		m, ok := isPossible(comb...)
		if !ok {
			continue
		}
		return m
	}
	panic("x")
}

func check(m, n []byte, idx1, idx2 int) bool {
	if m[idx1] != 0 && m[idx1] != n[idx2] {
		return false
	}
	m[idx1] = n[idx2]
	return true
}

func isPossible(known ...string) (string, bool) {
	m := make([]byte, 7)

	for _, k := range known {

		switch len(k) {
		case 2:
			if !check(m, []byte(k), 2, 0) {
				return "", false
			}
			if !check(m, []byte(k), 5, 1) {
				return "", false
			}
		case 4:
			if !check(m, []byte(k), 1, 0) {
				return "", false
			}
			if !check(m, []byte(k), 2, 1) {
				return "", false
			}
			if !check(m, []byte(k), 3, 2) {
				return "", false
			}
			if !check(m, []byte(k), 5, 3) {
				return "", false
			}
		case 3:
			if !check(m, []byte(k), 0, 0) {
				return "", false
			}
			if !check(m, []byte(k), 2, 1) {
				return "", false
			}
			if !check(m, []byte(k), 5, 2) {
				return "", false
			}
		case 7:
			if !check(m, []byte(k), 0, 0) {
				return "", false
			}
			if !check(m, []byte(k), 1, 1) {
				return "", false
			}
			if !check(m, []byte(k), 2, 2) {
				return "", false
			}
			if !check(m, []byte(k), 3, 3) {
				return "", false
			}
			if !check(m, []byte(k), 4, 4) {
				return "", false
			}
			if !check(m, []byte(k), 5, 5) {
				return "", false
			}
			if !check(m, []byte(k), 6, 6) {
				return "", false
			}
		}

	}
	return string(m), true
}

func Perm(a string, f func(string)) {
	perm([]byte(a), f, 0)
}

func perm(a []byte, f func(string), i int) {
	if i > len(a) {
		f(string(a))
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

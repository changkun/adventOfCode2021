package main

import (
	"math"
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

func Counting(fname string) int {
	s, _ := utils.Read(fname)

	var rate []int // counts 1
	total := 0
	for s.Scan() {
		tex := s.Text()
		total++
		ns := strings.Split(tex, "")
		if rate == nil {
			rate = make([]int, len(ns))
		}

		for i, n := range ns {
			nn, _ := strconv.Atoi(n)
			if nn == 1 {
				rate[i]++
			}
		}
	}

	gamma, epsilon := 0, 0
	for i, r := range rate {
		if r < total-r {
			rate[i] = 0
			epsilon += int(math.Pow(2, float64(len(rate)-i-1)))
		} else {
			rate[i] = 1
			gamma += int(math.Pow(2, float64(len(rate)-i-1)))
		}
	}

	return gamma * epsilon
}

type filterFunc func([]string, int, filterType) []string
type filterType func() (byte, byte)

func filter(all []string, idx int, ftype filterType) (ret []string) {
	counter := make([]int, len(all[0]))
	for _, s := range all {
		for i, n := range strings.Split(s, "") {
			if nn, _ := strconv.Atoi(n); nn == 1 {
				counter[i]++
			}
		}
	}
	a, b := ftype()
	for _, s := range all {
		if counter[idx] >= len(all)-counter[idx] {
			if s[idx] == a {
				ret = append(ret, s)
			}
		} else {
			if s[idx] == b {
				ret = append(ret, s)
			}
		}
	}
	return
}

func calc(all []string, filter filterFunc, ftype filterType) int {
	cur := make([]string, len(all))
	copy(cur, all)
	for i := 0; i < len(all[0]); i++ {
		if len(cur) == 1 {
			break
		}
		cur = filter(cur, i, ftype)
	}

	val := 0
	for i, n := range cur[0] {
		if n == '1' {
			val += int(math.Pow(2, float64(len(cur[0])-i-1)))
		}
	}
	return val
}

func LifeSupportRate(fname string) int {
	s, _ := utils.Read(fname)
	all := []string{}
	for s.Scan() {
		all = append(all, s.Text())
	}
	a := calc(all, filter, func() (byte, byte) { return '1', '0' })
	b := calc(all, filter, func() (byte, byte) { return '0', '1' })
	return a * b
}

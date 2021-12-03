package main_test

import (
	"bufio"
	"math"
	"strconv"
	"strings"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

func counting(t *testing.T, s *bufio.Scanner) int {
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
			nn, err := strconv.Atoi(n)
			if err != nil {
				t.Fatal(err)
			}
			if nn == 1 {
				rate[i]++
			}
		}
	}

	gamma := 0
	epsilon := 0
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

func TestPart1(t *testing.T) {
	s := utils.Read(t, "test.txt")
	if n := counting(t, s); n != 198 {
		t.Fatalf("want 198, got %v", n)
	}

	s = utils.Read(t, "input.txt")
	t.Log(counting(t, s))
}

func rate(t *testing.T, ss []string) ([]int, int) {
	counter := make([]int, len(ss[0]))
	total := 0
	for _, s := range ss {
		total++
		ns := strings.Split(s, "")
		for i, n := range ns {
			nn, err := strconv.Atoi(n)
			if err != nil {
				t.Fatal(err)
			}
			if nn == 1 {
				counter[i]++
			}
		}
	}
	return counter, total
}

func oxygenFilter(t *testing.T, ss []string, idx int) []string {
	counter, total := rate(t, ss)

	ret := []string{}
	for _, s := range ss {
		if counter[idx] >= total-counter[idx] { // keep strings with 1
			if s[idx] == '1' {
				ret = append(ret, s)
			}
		} else {
			if s[idx] == '0' {
				ret = append(ret, s)
			}
		}
	}
	return ret
}

func co2Filter(t *testing.T, ss []string, idx int) []string {
	counter, total := rate(t, ss)
	ret := []string{}

	// filtering idx
	for _, s := range ss {
		if counter[idx] >= total-counter[idx] {
			if s[idx] == '0' {
				ret = append(ret, s)
			}
		} else {
			if s[idx] == '1' {
				ret = append(ret, s)
			}
		}
	}
	return ret
}

func calc(t *testing.T, all []string, filter func(t *testing.T, ss []string, idx int) []string) int {
	cur := make([]string, len(all))
	copy(cur, all)
	for i := 0; i < len(all[0]); i++ {
		if len(cur) == 1 {
			break
		}
		cur = filter(t, cur, i)
	}

	ns := cur[0]
	val := 0
	for i, n := range ns {
		if n == '1' {
			val += int(math.Pow(2, float64(len(ns)-i-1)))
		}
	}
	return val
}

func lifeSupportRate(t *testing.T, fname string) int {
	s := utils.Read(t, fname)

	all := []string{}
	for s.Scan() {
		all = append(all, s.Text())
	}

	oxygen := calc(t, all, oxygenFilter)
	co2 := calc(t, all, co2Filter)
	return oxygen * co2
}

func TestPart2(t *testing.T) {

	n := lifeSupportRate(t, "test.txt")
	if n != 230 {
		t.Fatalf("want 230, got %v", n)
	}

	t.Log(lifeSupportRate(t, "input.txt"))
}

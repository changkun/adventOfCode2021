package main_test

import (
	"math"
	"strconv"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

func increase(t *testing.T, fname string) int {
	s := utils.Read(t, fname)
	last := math.MaxInt
	n := 0
	for s.Scan() {
		ns := s.Text()
		i, err := strconv.Atoi(ns)
		if err != nil {
			t.Fatalf("cannot convert text to number: %v, err: %v", ns, err)
		}

		if i > last {
			n++
		}
		last = i
	}
	return n
}

func TestPart1(t *testing.T) {
	n := increase(t, "test.txt")
	if n != 7 {
		t.Fatalf("test didn't pass, expect 7 got %v", n)
	}

	t.Log(increase(t, "input.txt"))
}

func increase2(t *testing.T, fname string) int {
	s := utils.Read(t, fname)
	nums := []int{}
	for s.Scan() {
		ns := s.Text()
		i, err := strconv.Atoi(ns)
		if err != nil {
			t.Fatalf("cannot convert text to number: %v, err: %v", ns, err)
		}
		nums = append(nums, i)
	}

	wins := []int{}
	for i := 0; i < len(nums)-2; i++ {
		wins = append(wins, nums[i]+nums[i+1]+nums[i+2])
	}

	last := math.MaxInt
	n := 0
	for _, sum := range wins {
		if sum > last {
			n++
		}
		last = sum
	}
	return n
}

func TestPart2(t *testing.T) {
	n := increase2(t, "test.txt")
	if n != 5 {
		t.Fatalf("test didn't pass, expect 5 got %v", n)
	}
	t.Log(increase2(t, "input.txt"))
}

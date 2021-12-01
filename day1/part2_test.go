package main_test

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"strconv"
	"testing"
)

func increase2(t *testing.T, fname string) int {
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
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

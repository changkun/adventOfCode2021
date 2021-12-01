package main_test

import (
	"bufio"
	"bytes"
	"math"
	"os"
	"strconv"
	"testing"
)

func increase(t *testing.T, fname string) int {
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
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

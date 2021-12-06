package main

import (
	"strconv"
	"strings"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

func readAllFishes(fname string) []int {
	s, _ := utils.Read(fname)
	fishes := []int{}
	for s.Scan() {
		tt := s.Text()
		ns := strings.Split(tt, ",")
		for _, n := range ns {
			nn, _ := strconv.Atoi(n)
			fishes = append(fishes, nn)
		}
	}
	return fishes
}

func TestPart1(t *testing.T) {
	fishes := readAllFishes("test.txt")
	if n := LanternfishGrow(fishes, 18); n != 26 {
		t.Fatalf("want 26 got %v", n)
	}
	if n := LanternfishGrow(fishes, 80); n != 5934 {
		t.Fatalf("want 5934 got %v", n)
	}

	fishes = readAllFishes("input.txt")
	t.Log(LanternfishGrow(fishes, 80))
}

func TestPart2(t *testing.T) {
	fishes := readAllFishes("test.txt")
	if n := LanternfishGrow(fishes, 256); n != 26984457539 {
		t.Fatalf("want 26984457539 got %v", n)
	}

	fishes = readAllFishes("input.txt")
	t.Log(LanternfishGrow(fishes, 256))
}

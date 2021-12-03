package main_test

import (
	"strconv"
	"strings"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

type pos struct {
	x, y int
}

func calpos(t *testing.T, fname string) pos {
	s := utils.Read(t, fname)
	p := pos{}

	for s.Scan() {
		cmds := strings.Split(s.Text(), " ")
		d, err := strconv.Atoi(cmds[1])
		if err != nil {
			t.Fatal(err)
		}
		switch cmds[0] {
		case "forward":
			p.x += d
		case "down":
			p.y += d
		case "up":
			p.y -= d
		}
	}
	return p
}

func TestPart1(t *testing.T) {
	if p := calpos(t, "test.txt"); p.x*p.y != 150 {
		t.Fatalf("test didn't pass, want 150, got %v", p.x*p.y)
	}

	p := calpos(t, "input.txt")

	t.Log(p.x * p.y)
}

func calpos2(t *testing.T, fname string) pos {
	s := utils.Read(t, fname)
	p := pos{}
	aim := 0

	for s.Scan() {
		cmds := strings.Split(s.Text(), " ")
		d, err := strconv.Atoi(cmds[1])
		if err != nil {
			t.Fatal(err)
		}
		switch cmds[0] {
		case "down":
			aim += d
		case "up":
			aim -= d
		case "forward":
			p.x += d
			p.y += aim * d
		}
	}
	return p
}

func TestPart2(t *testing.T) {
	p := calpos2(t, "test.txt")
	if p.x*p.y != 900 {
		t.Fatalf("test didn't pass, want 900, got %v", p.x*p.y)
	}
	t.Log(p)

	p = calpos2(t, "input.txt")
	t.Log(p.x * p.y)
}

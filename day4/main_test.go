package main

import "testing"

func TestPart1(t *testing.T) {
	m, d := computeWin("test.txt")
	v := m[0].SumPositive() * d[0]
	if v != 4512 {
		t.Fatalf("want 4512 got %v", v)
	}

	m, d = computeWin("input.txt")
	t.Log(m[0].SumPositive() * d[0])
}

func TestPart2(t *testing.T) {
	m, d := computeWin("test.txt")
	v := m[len(m)-1].SumPositive() * d[len(d)-1]
	if v != 1924 {
		t.Fatalf("want 1924 got %v", v)
	}

	m, d = computeWin("input.txt")
	t.Log(m[len(m)-1].SumPositive() * d[len(d)-1])
}

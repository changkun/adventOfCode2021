package main

import "testing"

func TestPart1(t *testing.T) {
	n := calc("test.txt", false)
	if n != 5 {
		t.Fatalf("want 5 got: %v", n)
	}

	t.Log(calc("input.txt", false))
}

func TestPart2(t *testing.T) {
	n := calc("test.txt", true)
	if n != 12 {
		t.Fatalf("want 12 got: %v", n)
	}

	t.Log(calc("input.txt", true))
}

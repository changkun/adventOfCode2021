package main

import "testing"

func TestPart1(t *testing.T) {
	n := Increase("test.txt", 1)
	if n != 7 {
		t.Fatalf("test didn't pass, expect 7 got %v", n)
	}

	t.Log(Increase("input.txt", 1))
}

func TestPart2(t *testing.T) {
	n := Increase("test.txt", 3)
	if n != 5 {
		t.Fatalf("test didn't pass, expect 5 got %v", n)
	}
	t.Log(Increase("input.txt", 3))
}

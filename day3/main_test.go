package main

import "testing"

func TestPart1(t *testing.T) {
	if n := Counting("test.txt"); n != 198 {
		t.Fatalf("want 198, got %v", n)
	}
	t.Log(Counting("input.txt"))
}

func TestPart2(t *testing.T) {

	n := LifeSupportRate("test.txt")
	if n != 230 {
		t.Fatalf("want 230, got %v", n)
	}

	t.Log(LifeSupportRate("input.txt"))
}

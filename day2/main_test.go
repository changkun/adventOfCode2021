package main

import "testing"

func TestPart1(t *testing.T) {
	if p := Process("test.txt", false); p.x*p.y != 150 {
		t.Fatalf("test didn't pass, want 150, got %v", p.x*p.y)
	}

	p := Process("input.txt", false)
	t.Log(p.x * p.y)
}

func TestPart2(t *testing.T) {
	p := Process("test.txt", true)
	if p.x*p.y != 900 {
		t.Fatalf("test didn't pass, want 900, got %v", p.x*p.y)
	}
	t.Log(p)

	p = Process("input.txt", true)
	t.Log(p.x * p.y)
}

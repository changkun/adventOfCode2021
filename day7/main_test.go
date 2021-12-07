package main

import "testing"

func TestPart1(t *testing.T) {
	t.Log(leatFuel(readPos("test.txt")))
	t.Log(leatFuel(readPos("input.txt")))
}

func TestPart2(t *testing.T) {
	t.Log(leatFuel2(readPos("test.txt")))
	t.Log(leatFuel2(readPos("input.txt")))
}

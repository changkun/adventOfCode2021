package main

import (
	"strconv"
	"strings"
	"testing"

	"changkun.de/x/adventOfCode2021/utils"
)

func TestPart1(t *testing.T) {
	t.Log(uniqueCount("test.txt"))
	t.Log(uniqueCount("input.txt"))
}

func TestPart2(t *testing.T) {
	s, _ := utils.Read("tiny.txt")
	for s.Scan() {
		tt := s.Text()
		patterns := strings.Split(tt, " | ")
		digitsBefore := strings.Split(patterns[0], " ")
		digitsAfter := strings.Split(patterns[1], " ")

		known := []string{}
		for _, d := range digitsAfter {
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				known = append(known, d)
			}
		}
		for _, d := range digitsBefore {
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				known = append(known, d)
			}
		}
		m := decode(known...)

		mm := [10][]string{
			0: {string([]byte{m[0], m[1], m[2], m[4], m[5], m[6]})},
			1: {string([]byte{m[2], m[5]})},
			2: {string([]byte{m[1], m[2], m[3], m[4], m[6]})},
			3: {string([]byte{m[0], m[2], m[3], m[5], m[6]})},
			4: {string([]byte{m[1], m[2], m[3], m[5]})},
			5: {string([]byte{m[0], m[1], m[3], m[5], m[6]})},
			6: {string([]byte{m[0], m[1], m[3], m[4], m[5], m[6]})},
			7: {string([]byte{m[0], m[2], m[5]})},
			8: {string([]byte{m[0], m[1], m[2], m[3], m[4], m[5], m[6]})},
			9: {string([]byte{m[0], m[1], m[2], m[3], m[5], m[6]})},
		}
		t.Log(m, mm, digitsAfter)

		for i := range mm {
			Perm(mm[i][0], func(s string) {
				mm[i] = append(mm[i], s)
			})
		}

		val := ""
		for _, d := range digitsAfter {
			dd := 0
			for i := 0; i < 10; i++ {
				for _, v := range mm[i] {
					if strings.Compare(v, d) == 0 {
						dd = i
						goto co
					}
				}
			}
		co:
			// 	if dd == -1 {
			// 		panic("impossible")
			// 	}
			val += strconv.Itoa(dd)
			t.Log(val, dd)
		}

		break
	}
}

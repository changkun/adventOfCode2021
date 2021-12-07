package main

import (
	"math"
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

func readPos(fname string) (ret []int) {
	s, _ := utils.Read(fname)

	for s.Scan() {
		tt := s.Text()
		ns := strings.Split(tt, ",")
		for _, n := range ns {
			nn, _ := strconv.Atoi(n)
			ret = append(ret, nn)
		}
	}
	return
}

func leatFuel(pos []int) int {
	sum := make([]int, len(pos))

	for j := range sum {
		for _, n := range pos {
			sum[j] += int(math.Abs(float64(n - pos[j])))
		}
	}

	min := math.MaxInt
	for j := range sum {
		if sum[j] < min {
			min = sum[j]
		}
	}
	return min
}

func minmax(ns ...int) (min, max int) {
	min = math.MaxInt
	max = math.MinInt
	for i := range ns {
		min = int(math.Min(float64(min), float64(ns[i])))
		max = int(math.Max(float64(max), float64(ns[i])))
	}
	return
}

func sum(ns ...int) (n int) {
	for i := range ns {
		n += ns[i]
	}
	return
}

func leatFuel2(pos []int) int {
	moveCost := func(pos []int, x int) int {
		costs := make([]int, len(pos))
		for i, n := range pos {
			lenth := int(math.Abs(float64(n - x)))
			cost := int(float64((lenth+1)*lenth) * 0.5)
			costs[i] = cost
		}
		sum := sum(costs...)
		return sum
	}

	min, max := minmax(pos...)
	allMoveCosts := []int{}
	for i := min; i <= max; i++ {
		allMoveCosts = append(allMoveCosts, moveCost(pos, i))
	}

	min, _ = minmax(allMoveCosts...)
	return min
}

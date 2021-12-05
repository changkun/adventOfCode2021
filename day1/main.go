package main

import (
	"math"
	"strconv"

	"changkun.de/x/adventOfCode2021/utils"
)

func readSlice(fname string) []int {
	s, _ := utils.Read(fname)
	nums := []int{}
	for s.Scan() {
		ns := s.Text()
		i, _ := strconv.Atoi(ns)
		nums = append(nums, i)
	}
	return nums
}

func calc(nums []int) int {
	last := math.MaxInt
	n := 0
	for _, sum := range nums {
		if sum > last {
			n++
		}
		last = sum
	}
	return n
}

func Increase(fname string, winSize int) int {
	if winSize == 0 {
		return 0
	}

	nums := readSlice(fname)
	for i := 0; i < len(nums)-winSize+1; i++ {
		sum := 0
		for j := 0; j < winSize; j++ {
			sum += nums[i+j]
		}
		nums[i] = sum
	}
	return calc(nums[:len(nums)-winSize+1])
}

package main

func LanternfishGrow(fishes []int, days int) int {
	a := map[int]int{}
	for _, v := range fishes {
		a[v]++
	}
	for i := 0; i < days; i++ {
		a[(i+7)%9] += a[i%9]
	}
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

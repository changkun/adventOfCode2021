package main

import (
	"log"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

type node struct {
	current string
	next    []*node
}

func build(fname string) map[string][]string {
	paths := map[string][]string{}
	s, _ := utils.Read(fname)
	for s.Scan() {
		tt := s.Text()
		edge := strings.Split(tt, "-")
		paths[edge[0]] = append(paths[edge[0]], edge[1])
		paths[edge[1]] = append(paths[edge[1]], edge[0])
	}
	return paths
}

func explore(fname string, part2 bool) int {
	paths := build(fname)

	var count func(curCave string, visited map[string]int, anyTwiceSmall bool) int
	count = func(curCave string, visited map[string]int, anyTwiceSmall bool) int {
		if curCave == "end" {
			return 1
		}

		res := 0
		for _, cave := range paths[curCave] {
			big := strings.Compare(strings.ToUpper(cave), cave) == 0
			small := !big && cave != "start" && cave != "end"
			_, seen := visited[cave]
			if big || !seen {
				visited[cave] = 0
				res += count(cave, visited, anyTwiceSmall)
				delete(visited, cave)
			} else if part2 && small && !anyTwiceSmall {
				res += count(cave, visited, true)
			}
		}
		return res
	}

	return count("start", map[string]int{"start": 0}, false)
}

func main() {
	log.Println(explore("test.txt", false))
	log.Println(explore("input.txt", false))
	log.Println(explore("test.txt", true))
	log.Println(explore("input.txt", true))
}

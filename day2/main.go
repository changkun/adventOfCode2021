package main

import (
	"strconv"
	"strings"

	"changkun.de/x/adventOfCode2021/utils"
)

type Submarine struct {
	x, y    int
	aim     int
	withAim bool
}

func newSubmarine(withAim bool) Submarine {
	return Submarine{withAim: withAim}
}

func (p *Submarine) forward(d int) {
	if !p.withAim {
		p.x += d
		return
	}
	p.x += d
	p.y += p.aim * d
}

func (p *Submarine) down(d int) {
	if !p.withAim {
		p.y += d
		return
	}
	p.aim += d
}

func (p *Submarine) up(d int) {
	if !p.withAim {
		p.y -= d
		return
	}
	p.aim -= d
}

func (p *Submarine) process(cmd string) {
	args := strings.Fields(cmd)
	d, _ := strconv.Atoi(args[1])
	switch args[0] {
	case "forward":
		p.forward(d)
	case "down":
		p.down(d)
	case "up":
		p.up(d)
	}
}

func Process(fname string, withAim bool) Submarine {
	s, _ := utils.Read(fname)
	p := newSubmarine(withAim)
	for s.Scan() {
		p.process(s.Text())
	}
	return p
}

package day2

import (
	"strconv"
	"strings"
)

type PositionWithAim struct {
	HorizontalPosition int
	Depth              int
	Aim                int
}

func (p *PositionWithAim) Move(command string) {
	commandSlice := strings.Split(command, " ")
	action := commandSlice[0]
	unit, _ := strconv.Atoi(commandSlice[1])

	switch action {
	case "forward":
		p.HorizontalPosition += unit
		p.Depth = p.Depth + (p.Aim * unit)
	case "up":
		p.Aim -= unit
	case "down":
		p.Aim += unit
	}
}

func (p *PositionWithAim) GetPosition() int {
	return p.HorizontalPosition * p.Depth
}

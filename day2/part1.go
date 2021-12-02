package day2

import (
	"strconv"
	"strings"
)

type Position struct {
	HorizontalPosition int
	Depth              int
}

func (p *Position) Move(command string) {
	commandSlice := strings.Split(command, " ")
	action := commandSlice[0]
	unit, _ := strconv.Atoi(commandSlice[1])

	switch action {
	case "forward":
		p.HorizontalPosition += unit
	case "up":
		p.Depth -= unit
	case "down":
		p.Depth += unit
	}
}

func (p *Position) GetPosition() int {
	return p.HorizontalPosition * p.Depth
}

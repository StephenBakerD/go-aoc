package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part2_example(t *testing.T) {
	p := PositionWithAim{HorizontalPosition: 0, Depth: 0, Aim: 0}
	p.Move("forward 5")
	p.Move("down 5")
	p.Move("forward 8")
	p.Move("up 3")
	p.Move("down 8")
	p.Move("forward 2")

	assert.Equal(t, 15, p.HorizontalPosition)
	assert.Equal(t, 60, p.Depth)
	assert.Equal(t, 900, p.GetPosition())
}

func Test_part2_input(t *testing.T) {
	p := PositionWithAim{HorizontalPosition: 0, Depth: 0, Aim: 0}
	input := readInput("input.txt")
	for _, command := range input {
		p.Move(command)
	}
	
	t.Logf("Horizaontal Position: %d, Depth: %d, Position %d", p.HorizontalPosition, p.Depth, p.GetPosition())
}

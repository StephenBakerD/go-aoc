package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1_example(t *testing.T) {
	p := Position{HorizontalPosition: 0, Depth: 0}
	p.Move("forward 5")
	p.Move("down 5")
	p.Move("forward 8")
	p.Move("up 3")
	p.Move("down 8")
	p.Move("forward 2")

	assert.Equal(t, 15, p.HorizontalPosition)
	assert.Equal(t, 10, p.Depth)
	assert.Equal(t, 150, p.GetPosition())
}

func Test_part1_input(t *testing.T) {
	p := Position{HorizontalPosition: 0, Depth: 0}
	input := readInput("input.txt")
	for _, command := range input {
		p.Move(command)
	}
	
	t.Logf("Horizaontal Position: %d, Depth: %d, Position %d", p.HorizontalPosition, p.Depth, p.GetPosition())
}

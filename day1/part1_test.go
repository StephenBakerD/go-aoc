package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1_example_input(t *testing.T) {
	answer := Part1(199, 200, 208, 210, 200, 207, 240, 269, 260, 263)
	assert.Equal(t, 7, answer)
}

func Test_part1_my_input(t *testing.T) {
	values := readInput("input.txt")
	answer := Part1(values...)
	t.Log(answer)
}

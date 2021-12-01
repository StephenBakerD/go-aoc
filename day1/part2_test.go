package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part2_example_input(t *testing.T) {
	actual :=  Part2(199,200,208,210,200,207,240,269,260,263)
	assert.Equal(t, 5, actual)
}

func Test_part2_my_input(t *testing.T) {
	input := readInput("input.txt")
	result :=  Part2(input...)
	t.Log(result)
}
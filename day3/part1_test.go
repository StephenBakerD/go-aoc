package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1_example(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}

	gammaString, gammaInt, epsilonString, epsilonInt := calculatePower(input)

	assert.Equal(t, gammaString, "10110")
	assert.Equal(t, epsilonString, "01001")
	assert.Equal(t, gammaInt, 22)
	assert.Equal(t, epsilonInt, 9)
}

func Test_part1_input(t *testing.T) {
	input := readInput("input.txt")

	_, gammaInt, _, epsilonInt := calculatePower(input)
	t.Log(gammaInt * epsilonInt)
}

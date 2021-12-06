package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part2_example(t *testing.T) {
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

	oxygenGeneratorRating, co2ScrubberRating := calculateLifeSupport(input)

	assert.Equal(t, oxygenGeneratorRating, 23)
	assert.Equal(t, co2ScrubberRating, 10)
}

func Test_part2_input(t *testing.T) {
	input := readInput("input.txt")

	oxygenGeneratorRating, co2ScrubberRating := calculateLifeSupport(input)
	t.Log(oxygenGeneratorRating * co2ScrubberRating)
}

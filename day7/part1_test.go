package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//acceptance test
func Test_part1_example(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	crabs, _ := parseInput(input)
	targetNumber, fuelCost, err := alignCrabs_part1(crabs)

	assert.Nil(t, err)
	assert.Equal(t, 2, targetNumber)
	assert.Equal(t, 37, fuelCost)
}

//acceptance test
func Test_part1(t *testing.T) {
	input := readFile("input.txt")
	crabs, _ := parseInput(input)
	targetNumber, fuelCost, err := alignCrabs_part1(crabs)
	
	assert.Nil(t, err)
	t.Logf("target: %v fuel cost: %v", targetNumber, fuelCost)
}

package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//acceptance test
func Test_part2_example(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	crabs, _ := parseInput(input)
	targetNumber, fuelCost, err := alignCrabs_part2(crabs)

	assert.Nil(t, err)
	assert.Equal(t, 5, targetNumber)
	assert.Equal(t, 168, fuelCost)
}

//acceptance test
func Test_part2(t *testing.T) {
	input := readFile("input.txt")
	crabs, _ := parseInput(input)
	targetNumber, fuelCost, err := alignCrabs_part2(crabs)
	
	assert.Nil(t, err)
	t.Logf("target: %v fuel cost: %v", targetNumber, fuelCost)
}


//unit test
func Test_calcFuelCostFromDelta(t *testing.T) {
	type args struct {
		delta int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "delta: 2", args: args{delta: 2}, want: 3},
		{name: "delta: 3", args: args{delta: 3}, want: 6},
		{name: "delta: 11", args: args{delta: 11}, want: 66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuelCostFromDelta(tt.args.delta); got != tt.want {
				t.Errorf("calcFuelCostFromDelta() = %v, want %v", got, tt.want)
			}
		})
	}
}

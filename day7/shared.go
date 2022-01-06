package day7

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) string {
	//read complete file into bytes
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	return string(bytes)
}

func parseInput(input string) ([]int, error) {
	stringsSlice := strings.Split(input, ",")
	intSlice := make([]int, len(stringsSlice))

	for i, stringVal := range stringsSlice {
		intVal, err := strconv.Atoi(stringVal)
		if err != nil {
			return nil, err
		}
		intSlice[i] = intVal
	}

	return intSlice, nil
}

func minMax(slice []int) (min int, max int, err error) {
	if len(slice) == 0 {
		err = errors.New("slice length cannot be zero")
		return
	}

	//initialize min and max values
	min = slice[0]
	max = slice[0]

	for _, val := range slice[1:] {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return
}

func buildFuelCostMap(crabs []int) (map[int]int, error) {
	result := make(map[int]int) //target number:fuel cost

	//add all possible whole numbers between min/max
	min, max, err := minMax(crabs)
	if err != nil {
		return nil, err
	}
	for i := min; i <= max; i++ {
		result[i] = 0
	}

	return result, nil
}

func calculateFuelCostsEscalated(crabs []int, possibleFuelCosts map[int]int) map[int]int {
	result := make(map[int]int, len(possibleFuelCosts))

	//calculate all possible fuel costs
	for target := range possibleFuelCosts {
		for _, crab := range crabs {
			delta := int(math.Abs(float64(target - crab))) //get absolute delta for crab moving to target
			result[target] += calcFuelCostFromDelta(delta)
		}
	}

	return result
}

func calculateFuelCosts(crabs []int, possibleFuelCosts map[int]int) map[int]int {
	result := make(map[int]int, len(possibleFuelCosts))

	//calculate all possible fuel costs
	for target := range possibleFuelCosts {
		for _, crab := range crabs {
			delta := int(math.Abs(float64(target - crab))) //get absolute delta for crab moving to target
			result[target] += delta
		}
	}

	return result
}

func getLowestFuelCost(possibleFuelCosts map[int]int) (int, int) {
	var minTarget int = 0
	var minFuelCost int = math.MaxInt
	for target, fuelCost := range possibleFuelCosts {
		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
			minTarget = target
		}
	}

	return minTarget, minFuelCost
}

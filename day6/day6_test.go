package day6

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1_example(t *testing.T) {
	input := "3,4,3,1,2"

	lanternfish := parseInput(input)
	amount := countLanternfish_part1(lanternfish, 80)

	assert.Equal(t, 5934, amount)
}

func Test_part1(t *testing.T) {
	input := readInput("input.txt")

	lanternfish := parseInput(input)
	amount := countLanternfish_part1(lanternfish, 80)

	t.Log(amount)
}

func Test_part2_example(t *testing.T) {
	input := "3,4,3,1,2"

	lanternfish := parseInput(input)
	amount := countLanternfish_part2(lanternfish, 256)

	assert.Equal(t, 26984457539, amount)
}

func Test_part2(t *testing.T) {
	input := readInput("input.txt")

	lanternfish := parseInput(input)
	amount := countLanternfish_part2(lanternfish, 256)

	t.Log(amount)
}

func parseInput(input string) []int {
	inputStrings := strings.Split(input, ",")
	inputInts := make([]int, len(inputStrings))

	//convert strings to ints
	for i, inputString := range inputStrings {
		if val, err := strconv.Atoi(inputString); err == nil {
			inputInts[i] = val
		} else {
			panic(err)
		}
	}

	return inputInts
}

func countLanternfish_part1(lanternFish []int, daysInFuture int) int {
	for i := 0; i < daysInFuture; i++ {
		lanternFishCopy := lanternFish[:] //copy slice in order to iterate and modify at the same time

		for j := range lanternFishCopy {
			if lanternFish[j] > 0 {
				lanternFish[j]--
			} else {
				lanternFish = append(lanternFish, 8)
				lanternFish[j] = 6
			}
		}
	}

	return len(lanternFish)
}

func countLanternfish_part2(lanternFish []int, daysInFuture int) int {
	var lanternFishMap = map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0} //day:count

	//add initial lantern fish values to map
	for _, f := range lanternFish {
		lanternFishMap[f]++
	}

	for i := 0; i < daysInFuture; i++ {
		//shift counts down by 1 day
		lanternFishMapCopy := copyMap(lanternFishMap)
		for j := 8; j > 0; j-- {
			lanternFishMap[j-1] = lanternFishMapCopy[j]
		}

		//create new fish for any fish at zero
		lanternFishMap[8] = lanternFishMapCopy[0]

		//reset count to 6 for any fish at zero
		lanternFishMap[6] += lanternFishMapCopy[0]
	}

	//calculate sum of all lantern fish
	sum := 0
	for _, v := range lanternFishMap {
		sum += v
	}

	return sum
}

func copyMap(originalMap map[int]int) map[int]int {
	copy := make(map[int]int)

	for k, v := range originalMap {
		copy[k] = v
	}

	return copy
}

func readInput(fileName string) string {
	//read complete file into bytes
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	return string(bytes)
}

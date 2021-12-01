package day1

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1_example_input(t *testing.T) {
	answer := Part1(199, 200, 208, 210, 200, 207, 240, 269, 260, 263)
	assert.Equal(t, 7, answer)
}

func Test_part1_my_input(t *testing.T) {
	values := readInput()
	answer := Part1(values...)
	t.Log(answer)
}

func readInput() []int {
	//read complete file into bytes
	bytes, err := os.ReadFile("part1_input.txt")
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	content := string(bytes)

	//split lines, and convert to int[]
	var lines []int
	for _, line := range strings.Split(content, "\r\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		lines = append(lines, val)
	}

	return lines
}

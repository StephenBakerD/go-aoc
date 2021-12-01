package day1

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day1_Example_Input(t *testing.T) {
	answer := Count(199, 200, 208, 210, 200, 207, 240, 269, 260, 263)
	assert.Equal(t, 7, answer)
}

func Test_Day1_My_Input(t *testing.T) {
	values := readInput()
	answer := Count(values...)
	t.Log(answer)
}

func readInput() []int {
	//read complete file into bytes
	bytes, err := os.ReadFile("input.txt")
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

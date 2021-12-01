package day1

import (
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) []int {
	//read complete file into bytes
	bytes, err := os.ReadFile(fileName)
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

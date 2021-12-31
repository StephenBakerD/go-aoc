package day4

import (
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) (*[]int, *[][]string) {
	//read complete file into bytes
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	content := string(bytes)

	lines := strings.Split(content, "\r\n")

	var randomNumbers []int
	for _, val := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(val)
		randomNumbers = append(randomNumbers, i)
	}

	var boards [][]string
	var board []string
	lines = lines[2:]
	for i, line := range lines {
		if line == "" {
			if i != 0 {
				boards = append(boards, board)
				board = make([]string, 0)
			}
		} else {
			board = append(board, line)
		}
	}

	if len(board) > 0 {
		boards = append(boards, board)
	}

	return &randomNumbers, &boards
}

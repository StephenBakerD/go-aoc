package day4

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1_example(t *testing.T) {
	randomNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	var boards []Board
	boards = append(boards, MakeBoard(
		"22 13 17 11 0",
		"8 2 23 4 24",
		"21 9 14 16 7",
		"6 10 3 18 5",
		"1 12 20 15 19",
	))
	boards = append(boards, MakeBoard(
		"14 21 17 24 4",
		"10 16 15 9 19",
		"18 8 23 26 20",
		"22 11 13 6 5",
		"2 0 12 3 7",
	))

	winningBoard, winningNumber := CallNumbersForFirstToWin(&boards, &randomNumbers)
	score := CalculateScore(winningBoard, winningNumber)
	assert.Equal(t, 4512, score)
}

func Test_part1_input(t *testing.T) {
	randomNumbers, allBoards := readInput("input.txt")
	var boards []Board
	for _, boardStringSlice := range *allBoards {
		boards = append(boards, MakeBoard(boardStringSlice...))
	}
	
	//verify boards were parsed correctly
	for _, board := range boards {
		for _, col := range board.Columns {
			if len(col) < 5 {
				assert.Error(t, errors.New("column length less than 5"))
			}
		}
		for _, row := range board.Rows {
			if len(row) < 5 {
				assert.Error(t, errors.New("row length less than 5"))
			}
		}
	}
	
	winningBoard, winningNumber := CallNumbersForFirstToWin(&boards, randomNumbers)
	score := CalculateScore(winningBoard, winningNumber)
	t.Log(score)
}

func Test_part2_example(t *testing.T) {
	randomNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	var boards []Board
	boards = append(boards, MakeBoard(
		"22 13 17 11 0",
		"8 2 23 4 24",
		"21 9 14 16 7",
		"6 10 3 18 5",
		"1 12 20 15 19",
	))
	boards = append(boards, MakeBoard(
		"3 15 0 2 22",
		"9 18 13 17 5",
		"19 8 7 25 23",
		"20 11 10 24 4",
		"14 21 16 12 6",
	))
	boards = append(boards, MakeBoard(
		"14 21 17 24 4",
		"10 16 15 9 19",
		"18 8 23 26 20",
		"22 11 13 6 5",
		"2 0 12 3 7",
	))

	lastBoardToWin, winningNumber := CallNumbersForLastToWin(&boards, &randomNumbers)
	score := CalculateScore(lastBoardToWin, winningNumber)
	assert.Equal(t, 1924, score)
}

func Test_part2_input(t *testing.T) {
	randomNumbers, allBoards := readInput("input.txt")
	var boards []Board
	for _, boardStringSlice := range *allBoards {
		boards = append(boards, MakeBoard(boardStringSlice...))
	}
	
	//verify boards were parsed correctly
	for _, board := range boards {
		for _, col := range board.Columns {
			if len(col) < 5 {
				assert.Error(t, errors.New("column length less than 5"))
			}
		}
		for _, row := range board.Rows {
			if len(row) < 5 {
				assert.Error(t, errors.New("row length less than 5"))
			}
		}
	}
	
	lastBoardToWin, winningNumber := CallNumbersForLastToWin(&boards, randomNumbers)
	score := CalculateScore(lastBoardToWin, winningNumber)
	t.Log(score)
}
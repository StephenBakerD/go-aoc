package day4

import (
	"strconv"
	"strings"
)

type Board struct {
	Squares []*Square
	Rows    map[int][]*Square
	Columns map[int][]*Square
}

type Square struct {
	Row       int
	Column    int
	Value     int
	WasCalled bool
}

func MakeBoard(rows ...string) Board {
	var result Board
	result.Columns = make(map[int][]*Square)
	result.Rows = make(map[int][]*Square)

	for irow, row := range rows {
		//clean up strings
		row = strings.Trim(row, " ")
		row = strings.Replace(row, "  ", " ", -1)

		for icol, col := range strings.Split(row, " ") {
			value, _ := strconv.Atoi(col)                                //parse string
			square := Square{Row: irow, Column: icol, Value: value}      //build square
			result.Columns[icol] = append(result.Columns[icol], &square) //add square to column slice
			result.Rows[irow] = append(result.Rows[irow], &square)       //add square to row slice
			result.Squares = append(result.Squares, &square)             //add square to all slice
		}
	}

	return result
}

func CallNumbersForFirstToWin(boards *[]Board, randomNumbers *[]int) (*Board, int) {
	for _, number := range *randomNumbers { //loop random numbers
		for _, board := range *boards { //loop boards
			for _, square := range board.Squares { //loop squares
				if square.Value == number {
					square.WasCalled = true
				}
			}

			if checkForWinner(board) {
				return &board, number
			}
		}
	}

	return nil, 0
}

func CallNumbersForLastToWin(boards *[]Board, randomNumbers *[]int) (*Board, int) {
	for _, number := range *randomNumbers { //loop random numbers
		var boardsToRemove []int

		for boardIndex, board := range *boards { //loop boards
			for _, square := range board.Squares { //loop squares
				if square.Value == number {
					square.WasCalled = true
				}
			}

			if checkForWinner(board) {
				if len(*boards) == 1 { //if only one board left, return it
					return &(*boards)[0], number
				} else { //otherwise mark the board to be removed from the slice
					boardsToRemove = append(boardsToRemove, boardIndex)
				}
			}
		}

		boards = removeBoards(boards, boardsToRemove...)
	}

	return nil, 0
}

func removeBoards(boards *[]Board, indexesToRemove ...int) *[]Board {
	//if nothing to remove, short circuit
	if len(indexesToRemove) == 0 {
		return boards
	}

	var result []Board

	for i, board := range *boards { //loop through boards
		var discard bool = false
		for _, indexToRemove := range indexesToRemove { //loop through indexesToRemove
			if i == indexToRemove {
				discard = true
			}
		}

		if !discard {
			result = append(result, board)
		}
	}

	return &result
}

func CalculateScore(board *Board, winningNumber int) int {
	unmarkedSum := 0

	for _, square := range board.Squares {
		if !square.WasCalled {
			unmarkedSum += square.Value
		}
	}

	return unmarkedSum * winningNumber
}

func checkForWinner(board Board) bool {
	for _, column := range board.Columns {
		var allCalled bool = true
		for _, square := range column {
			if !square.WasCalled {
				allCalled = false
				break
			}
		}

		if allCalled {
			return true
		}
	}

	for _, row := range board.Rows {
		var allCalled bool = true
		for _, square := range row {
			if !square.WasCalled {
				allCalled = false
				break
			}
		}

		if allCalled {
			return true
		}
	}

	return false
}

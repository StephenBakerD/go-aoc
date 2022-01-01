package day5

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parse_lines(t *testing.T) {
	inputRaw := "" +
		"2,0 -> 0,0" + "\n" +
		"0,0 -> 0,1"
	input := strings.Split(inputRaw, "\n")

	lines := parseLines(&input)

	line0 := (*lines)[0]                  //(0,0 -> 2,0)
	assert.Equal(t, 0, line0.Points[0].X) //0,0
	assert.Equal(t, 0, line0.Points[0].Y)
	assert.Equal(t, 1, line0.Points[1].X) //1, 0
	assert.Equal(t, 0, line0.Points[1].Y)
	assert.Equal(t, 2, line0.Points[2].X) //2, 0
	assert.Equal(t, 0, line0.Points[2].Y)

	line1 := (*lines)[1]                  //(0,0 -> 0,1)
	assert.Equal(t, 0, line1.Points[0].X) //0,0
	assert.Equal(t, 0, line1.Points[0].Y)
	assert.Equal(t, 0, line1.Points[1].X) //0,1
	assert.Equal(t, 1, line1.Points[1].Y)
}

func Test_part1_example(t *testing.T) {
	inputRaw := "" +
		"0,9 -> 5,9" + "\n" + //same y
		"8,0 -> 0,8" + "\n" +
		"9,4 -> 3,4" + "\n" + //same y
		"2,2 -> 2,1" + "\n" + //same x
		"7,0 -> 7,4" + "\n" + //same x
		"6,4 -> 2,0" + "\n" +
		"0,9 -> 2,9" + "\n" + //same y
		"3,4 -> 1,4" + "\n" + //same y
		"0,0 -> 8,8" + "\n" +
		"5,5 -> 8,2"
	input := strings.Split(inputRaw, "\n")

	lines := parseLines(&input)
	//there should only be 6 lines, since we only consider horizontal or vertical lines
	assert.Equal(t, 6, len(*lines))
	score := calculateOverlappingPoints(lines)

	assert.Equal(t, 5, score)
}

func Test_part1(t *testing.T) {
	input := readInput("input.txt")
	lines := parseLines(input)
	score := calculateOverlappingPoints(lines)

	t.Log(score)
}

func calculateOverlappingPoints(lines *[]Line) int {
	var points map[Point]int = make(map[Point]int)
	var score int

	for _, line := range *lines {
		for _, point := range line.Points {
			points[point]++ //if it doesn't exist, val will be 0
		}
	}

	for _, value := range points {
		if value > 1 {
			score++
		}
	}

	return score
}

func parseLines(input *[]string) *[]Line {
	var lines []Line
	for _, row := range *input {
		//clean up string
		row = strings.ReplaceAll(strings.Trim(row, "\t"), " ", "")

		columns := strings.Split(row, "->")
		if len(columns) != 2 {
			panic("input not in expected format. each line of input must contain two columns seperated by symbol ->")
		}
		//parse "x,y" into ["x", "y"]
		point1Strings := strings.Split(columns[0], ",")
		point2Strings := strings.Split(columns[1], ",")

		//parse ["x", "y"] into x = int & y = int
		point1X, err := strconv.Atoi(point1Strings[0])
		if err != nil {
			panic(err)
		}
		point1Y, err := strconv.Atoi(point1Strings[1])
		if err != nil {
			panic(err)
		}
		point2X, err := strconv.Atoi(point2Strings[0])
		if err != nil {
			panic(err)
		}
		point2Y, err := strconv.Atoi(point2Strings[1])
		if err != nil {
			panic(err)
		}

		if point1X != point2X && point1Y == point2Y { //0, 0 -> 5, 0 (would draw 6 points between)
			var points []Point
			//ensure point1X is always smallest
			if point1X > point2X {
				temp := point1X
				point1X = point2X
				point2X = temp
			}

			for x := point1X; x <= point2X; x++ {
				points = append(points, Point{X: x, Y: point1Y})
			}
			lines = append(lines, Line{Name: row, Points: points})
		} else if point1X == point2X && point1Y != point2Y { //0, 0 -> 0, 2 (would draw 3 points between)
			var points []Point

			//ensure point1Y is always smallest
			if point1Y > point2Y {
				temp := point1Y
				point1Y = point2Y
				point2Y = temp
			}

			for y := point1Y; y <= point2Y; y++ {
				points = append(points, Point{X: point1X, Y: y})
			}
			lines = append(lines, Line{Name: row, Points: points})
		}
	}

	return &lines
}

type Line struct {
	Name   string
	Points []Point
}

type Point struct {
	X int
	Y int
}

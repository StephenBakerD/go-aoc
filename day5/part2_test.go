package day5

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part2_example(t *testing.T) {
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

	lines := parseLines_part2(input)
	//diagonal lines now permitted
	assert.Equal(t, 10, len(lines))
	score := calculateOverlappingPoints(lines)

	assert.Equal(t, 12, score)
}

func Test_part2(t *testing.T) {
	input := readInput("input.txt")
	lines := parseLines_part2(input)
	score := calculateOverlappingPoints(lines)

	t.Log(score)
}

func parseLines_part2(input []string) []Line {
	var lines []Line
	for _, row := range input {
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

		point1 := Point{X: point1X, Y: point1Y}
		point2 := Point{X: point2X, Y: point2Y}

		slope, undefinedSlope := calculateSlope(point1, point2)

		if slope == 0 || undefinedSlope { //horizontal or vertical line
			points := interpolateHorizontalVerticalPoints(point1, point2)
			lines = append(lines, Line{Name: row, Points: points})
		} else if math.Abs(slope) == 1 { //diagonal line
			points := interpolateDiagonalPoints(point1, point2)
			lines = append(lines, Line{Name: row, Points: points})
		} //else discard line
	}

	return lines
}

func calculateSlope(point1 Point, point2 Point) (float64, bool) {
	var slope float64
	undefinedSlope := false

	xDelta := float64(point1.X - point2.X)
	yDelta := float64(point1.Y - point2.Y)

	if xDelta != 0 { //if not a vertical line with an undefined slope (since we cannot divide by zero)
		slope = yDelta / xDelta //slope = rise/run
	} else {
		undefinedSlope = true
	}

	return slope, undefinedSlope
}

func interpolateHorizontalVerticalPoints(point1 Point, point2 Point) []Point {
	var points []Point

	//ensure point1 is smaller
	if point1.X > point2.X || point1.Y > point2.Y {
		temp := point1
		point1 = point2
		point2 = temp
	}

	if point1.X != point2.X { //0,0 -> 5,0 (would draw 6 points between)
		for x := point1.X; x <= point2.X; x++ {
			points = append(points, Point{X: x, Y: point1.Y})
		}
	} else { //0,0 -> 0,2 (would draw 3 points between)
		for y := point1.Y; y <= point2.Y; y++ {
			points = append(points, Point{X: point1.X, Y: y})
		}
	}

	return points
}

func interpolateDiagonalPoints(point1 Point, point2 Point) []Point {
	var points []Point

	tempPoint := point1

	//figure out if we need to increment or decrement x
	var xIncrementer int
	if point2.X-point1.X > 0 {
		xIncrementer = 1
	} else {
		xIncrementer = -1
	}

	//figure out if we need to increment or decrement y
	var yIncrementer int
	if point2.Y-point1.Y > 0 {
		yIncrementer = 1
	} else {
		yIncrementer = -1
	}

	//interpolate all points between point1 and point 2
	points = append(points, point1)
	for tempPoint != point2 {
		tempPoint = Point{
			X: tempPoint.X + xIncrementer,
			Y: tempPoint.Y + yIncrementer}
		points = append(points, tempPoint)
	}

	return points
}

package day5

import (
	"os"
	"strings"
)

type Line struct {
	Name   string
	Points []Point
}

type Point struct {
	X int
	Y int
}

func readInput(fileName string) *[]string {
	//read complete file into bytes
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//convert bytes to string
	content := string(bytes)

	lines := strings.Split(content, "\r\n")
	return &lines
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

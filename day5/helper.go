package day5

import (
	"os"
	"strings"
)

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

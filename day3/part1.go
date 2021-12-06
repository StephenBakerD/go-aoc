package day3

import "strconv"

func calculatePower(input []string) (string, int, string, int) {
	var gammaString string
	var epsilonString string
	var counters []int

	//loop through each line
	for _, line := range input {
		//loop through each character
		for i, character := range line {
			//re-size counters if needed
			if len(counters) < i+1 {
				counters = append(counters, 0)
			}

			//subtract 1 if value is zero, add 1 if value is one
			if character == '0' {
				counters[i] -= 1
			} else if character == '1' {
				counters[i] += 1
			}
		}
	}

	//build gamma and epislon binary strings
	for _, counter := range counters {
		if counter > 0 {
			gammaString += "1"
			epsilonString += "0"
		} else if counter < 0 {
			gammaString += "0"
			epsilonString += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gammaString, 2, 32)
	epsilonInt, _ := strconv.ParseInt(epsilonString, 2, 32)

	return gammaString, int(gammaInt), epsilonString, int(epsilonInt)
}

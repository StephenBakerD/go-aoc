package day3

import "strconv"

func calculateLifeSupport(input []string) (int, int) {
	var oxygenGeneratorRating int64
	var co2ScrubberRating int64
	filter := input

	for i := 0; i < len(input[0]); i++ {
		counter := count(filter, i)
		filter = filterInputOxygen(filter, i, counter)
		if len(filter) == 1 {
			oxygenGeneratorRating, _ = strconv.ParseInt(filter[0], 2, 32)
			break
		}
	}

	filter = input
	for i := 0; i < len(input[0]); i++ {
		counter := count(filter, i)
		filter = filterInputCO2(filter, i, counter)
		if len(filter) == 1 {
			co2ScrubberRating, _ = strconv.ParseInt(filter[0], 2, 32)
			break
		}
	}

	return int(oxygenGeneratorRating), int(co2ScrubberRating)
}

func count(input []string, characterPosition int) int {
	var counter int

	for _, line := range input {
		if line[characterPosition] == '0' {
			counter -= 1
		} else if line[characterPosition] == '1' {
			counter += 1
		}
	}

	return counter
}

func filterInputOxygen(input []string, characterPosition int, counter int) []string {
	var result []string
	
	if counter >= 0 {
		for _, line := range input {
			if line[characterPosition] == '1' {
				result = append(result, line)
			}
		}
	} else if counter < 0 {
		for _, line := range input {
			if line[characterPosition] == '0' {
				result = append(result, line)
			}
		}
	}

	return result
}

func filterInputCO2(input []string, characterPosition int, counter int) []string {
	var result []string
	
	if counter >= 0 {
		for _, line := range input {
			if line[characterPosition] == '0' {
				result = append(result, line)
			}
		}
	} else if counter < 0 {
		for _, line := range input {
			if line[characterPosition] == '1' {
				result = append(result, line)
			}
		}
	}

	return result
}
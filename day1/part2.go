package day1

func Part2(measurements ...int) int {
	results := 0
	var groupsOfThree []int

	//build slice with sum of groups of 3
	for i, _ := range measurements {
		if i+2 <= len(measurements)-1 {
			sum := measurements[i] + measurements[i+1] + measurements[i+2]
			groupsOfThree = append(groupsOfThree, sum)
		}
	}

	//count how many times each element is greater than the previous
	previousVal := groupsOfThree[0]
	for _, val := range groupsOfThree[1:] {
		if val > previousVal {
			results++
		}
		previousVal = val
	}

	return results
}

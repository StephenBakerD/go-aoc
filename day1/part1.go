package day1

func Part1(measurements ...int) int{
	result := 0

	previousMeasurement := measurements[0]
	for _, i := range measurements[1:] {
		if i > previousMeasurement {
			result++;
		}
		previousMeasurement = i
	}
	
	return result
}
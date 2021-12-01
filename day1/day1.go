package day1

func Count(measurements ...int) int{
	result := 0

	previousMeasurement := measurements[0]
	for _, i := range measurements[1:] {
		if i > previousMeasurement {
			result++;
		}
		previousMeasurement = i
	}
	
	return result
}``
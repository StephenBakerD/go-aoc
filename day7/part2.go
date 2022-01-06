package day7

func alignCrabs_part2(crabs []int) (int, int, error) {
	fuelCostMap, err := buildFuelCostMap(crabs)
	if err != nil {
		return 0, 0, err
	}
	fuelCostMap = calculateFuelCostsEscalated(crabs, fuelCostMap)
	target, fuelCost := getLowestFuelCost(fuelCostMap)
	return target, fuelCost, nil
}

func calcFuelCostFromDelta(delta int) int {
	totalFuelCost := 0
	lastFuelCost := 1
	for i := delta; i > 0; i-- {
		totalFuelCost += lastFuelCost //add to sum
		lastFuelCost++                //increase the rate of fuel cost
	}

	return totalFuelCost
}

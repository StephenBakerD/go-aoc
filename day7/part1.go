package day7

func alignCrabs_part1(crabs []int) (int, int, error) {
	fuelCostMap, err := buildFuelCostMap(crabs)
	if err != nil {
		return 0, 0, err
	}
	fuelCostMap = calculateFuelCosts(crabs, fuelCostMap)
	target, fuelCost := getLowestFuelCost(fuelCostMap)
	return target, fuelCost, nil
}

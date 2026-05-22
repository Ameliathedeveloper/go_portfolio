package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	successfulCarsPerHour := float64(productionRate) * (successRate / 100)

	return successfulCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successfulCarsPerHour := float64(productionRate) * (successRate / 100)
	return int(successfulCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	individualCarCost := 10000
	batchOf10Cost := 95000

	if carsCount < 10 {
		return uint(carsCount * individualCarCost)
	}
	if carsCount%10 == 0 {
		return uint((carsCount / 10) * batchOf10Cost)
	}
	return uint((carsCount/10)*batchOf10Cost + (carsCount%10)*individualCarCost)

}
package lasagna

// OvenTime is the expected time in minutes for the lasagna to be ready.
// It does not change so it is declared as a constant.
const OvenTime = 40

// RemainingOvenTime takes a parameter of actualMinutesInOven and returns an int
// function return OvenTime (40 mins) minus the actualMinutesInOven to calculate the remaining time in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {

	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

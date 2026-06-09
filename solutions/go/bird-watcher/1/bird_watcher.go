package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0 // Initialize the bird count to zero

	for i := 0; i < len(birdsPerDay); i++ { // Loop through each day's bird count
		birdCount += birdsPerDay[i] // Add the current day's bird count to the total
	}
	return birdCount // Return the total bird count after summing all days
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	TotalBirdCount := 0 // Initialize the total bird count for the specified week to zero

	for i := 0; i < len(birdsPerDay); i++ { // Loop through each day's bird count
		if i >= (week-1)*7 && i < week*7 { // Check if the current index falls within the specified week
			TotalBirdCount += birdsPerDay[i] // Add the current day's bird count to the total if it belongs to the specified week
		}
		if i >= week*7 { // If the current index exceeds the last day of the specified week, break the loop
			break
		}
	}
	return TotalBirdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ { // Loop through each day's bird count
		if i%2 == 0 { // Check if the current index is even (alternate days)
			birdsPerDay[i]++ // Increment the bird count for the current day
		}
	}
	return birdsPerDay // Return the corrected bird counts
}

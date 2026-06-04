package speed

// define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,          // full battery
		batteryDrain: batteryDrain, // set the battery drain per drive
		speed:        speed,        // set the speed of the car
		distance:     0,            // initial distance is 0
	}
}

// define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance, // set the distance of the track
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	//if car battery is less than the battery drain, the car cannot drive and will not move
	if car.battery < car.batteryDrain {

		return car // return the car without changing its state
	}

	return Car{
		battery:      car.battery - car.batteryDrain, // reduce the battery by the battery drain
		batteryDrain: car.batteryDrain,               // keep the same battery drain
		speed:        car.speed,                      // keep the same speed
		distance:     car.distance + car.speed,       // increase the distance by the speed
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// if the car has enough battery and has driven the same distance as the track, it can finish the track
	if car.battery >= 0 && car.distance == track.distance {
		return true
	}
	// if the car has no battery and has not driven the same distance as the track, it cannot finish the track
	if car.battery <= 0 && car.distance < track.distance {
		return false
	}
	if car.battery > 0 && car.distance < track.distance { // if the car has enough battery but has not driven the same distance as the track, it can continue driving
		return CanFinish(Drive(car), track) // recursively drive the car until it can finish the track or runs out of battery
	}
	if car.battery < car.distance && car.distance < track.distance { // if the car has less battery than the distance it needs to drive and has not driven the same distance as the track, it cannot finish the track
		return false // otherwise, the car cannot finish the track

	} else {
		return CanFinish(Drive(car), track) // recursively drive the car until it can finish the track or runs out of battery
	}
}

// Package weather allows the contents of weather_forecast.go and weather_forecast_test.go to be tested for the presence of comments that follow Go's documentation conventions.
package weather

var (
	// CurrentCondition holds the current weather condition for the location specified in CurrentLocation.
	CurrentCondition string
	// CurrentLocation holds the name of the location for which the current weather condition is specified in CurrentCondition.
	CurrentLocation string
)

// Forecast returns a string describing the current weather condition for the specified city and condition. It also updates the global variables CurrentLocation and CurrentCondition with the provided city and condition values.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

// Package weather provides a simple weather forecast system.
package weather

var (
	// CurrentCondition holds the current weather condition.
	CurrentCondition string
	// CurrentLocation holds the location of the weather forecast.
	CurrentLocation string
)

// Forecast returns a string describing the current weather condition in a specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

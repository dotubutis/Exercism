// Package weather provides a function to format a string that reports the current weather condition fot a given location.  
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string 
// CurrentLocation represents the location for which the weather will be reported.
var CurrentLocation string

// Forecast returns a formatted string for reporting the weather based on the provided city and the weather conditon.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

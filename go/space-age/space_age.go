package space

import "math"

// Planet is a string corresponding to the name of one of the eight planets
type Planet string

// Age caculates the hypothetical age a person that has lived s seconds on earth would have on planet p.
func Age(s float64, p Planet) float64 {

	switch p {
	case "Earth":
		return 1.0000000 * s / 31557600
	case "Mercury":
		return (1 / 0.2408467) * s / 31557600
	case "Venus":
		return (1 / 0.61519726) * s / 31557600
	case "Mars":
		return (1 / 1.8808158) * s / 31557600
	case "Jupiter":
		return (1 / 11.862615) * s / 31557600
	case "Saturn":
		return (1 / 29.447498) * s / 31557600
	case "Uranus":
		return (1 / 84.016846) * s / 31557600
	case "Neptune":
		return (1 / 164.79132) * s / 31557600

	}
	return math.NaN()
}

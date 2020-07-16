// Package leap is a solution to the leap year problem of the side exercises of the Exercism go track
package leap

// IsLeapYear determines whether the year y is a leap year or not
func IsLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%100 == 0 {
		return false
	}
	if y%4 == 0 {
		return true
	}
	return false
}

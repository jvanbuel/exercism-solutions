// Package reverse is a solution to the reverse problem in the Exercism go track
package reverse

//Reverse reverse a string s
func Reverse(s string) string {
	var r string
	for _, c := range s {
		r = string(c) + r
	}
	return r
}

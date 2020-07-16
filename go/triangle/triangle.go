// Package triangle should have a package comment that summarizes what it's about.
package triangle

import "math"

// Kind is a kind of triangle
type Kind string

// Different kind of triangles
const (
	NaT = "Not a Triangle"
	Equ = "Equilateral"
	Iso = "Isocles"
	Sca = "Scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if (a*b*c <= 0) || math.IsNaN(a*b*c) || math.IsInf(a*b*c, 0) {
		return NaT
	}
	if (a+b < c) || (b+c < a) || (a+c < b) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

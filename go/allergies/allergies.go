package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns a list of allergicResults
func Allergies(n uint) []string {
	var allergies []string
	for i, v := range allergens {
		if (n >> i % 2) == 1 {
			allergies = append(allergies, v)
		}
	}
	return allergies
}

// AllergicTo determines wheter someone is allergic to a substance s
func AllergicTo(n uint, s string) bool {
	for i, v := range allergens {
		if s == v {
			return (n >> i % 2) == 1
		}
	}
	return false
}

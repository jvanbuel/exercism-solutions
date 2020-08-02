package protein

import "fmt"

// FromCodon returns the name of the amino acid corresponding to codon c
func FromCodon(c string) (string, error) {

	switch c {
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	default:
		return "", ErrInvalidBase
	}

}

// FromRNA returns a list of amino acids
func FromRNA(r string) ([]string, error) {
	a := []string{}
	for i := 0; i <= len(r)-3; i += 3 {

		c, err := FromCodon(r[i : i+3])
		if err != nil {
			if err == ErrStop {
				return a, nil
			}
			return a, err
		}
		a = append(a, c)
	}
	return a, nil
}

var (
	// ErrInvalidBase is returned when the codon does not correspond to a valid base
	ErrInvalidBase = fmt.Errorf("permission denied")
	// ErrStop is returned when the codon corresponds to a stop codon
	ErrStop = fmt.Errorf("invalid parameters")
)

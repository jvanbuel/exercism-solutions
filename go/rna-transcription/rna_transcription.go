package strand

// ToRNA returns the matching RNA complement for a dna string
func ToRNA(dna string) string {

	DNAToRNA := map[string]string{
		"A": "U",
		"T": "A",
		"G": "C",
		"C": "G",
	}
	var rna string
	for _, a := range dna {
		rna += DNAToRNA[string(a)]
	}
	return rna
}

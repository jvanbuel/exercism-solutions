// Package bob should have a package comment that summarizes what it's about.
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := strings.TrimSpace(remark)
	switch r {
	case "":
		return "Fine. Be that way!"
	case strings.ToLower(r):
		if r[len(r)-1] == '?' {
			return "Sure."
		}
		return "Whatever."
	case strings.ToUpper(r):
		if r[len(r)-1] == '?' {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	default:
		if r[len(r)-1] == '?' {
			return "Sure."
		}
		return "Whatever."
	}

}

package spanish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step0 is the removal of sttached pronouns
//
func step0(word *snowballword.SnowballWord) bool {

	// Search for the longest among the following suffixes 
	suffix1, suffix1Runes := word.FirstSuffix(
		"selas", "selos", "sela", "selo", "las", "les",
		"los", "nos", "me", "se", "la", "le", "lo",
	)

	// If the suffix empty or not in RV, we have nothing to do.
	if suffix1 == "" || word.RVstart > len(word.RS)-len(suffix1Runes) {
		return false
	}

	// We'll remove suffix1, if comes after one of the following
	suffix2, suffix2Runes := word.FirstSuffixAt(len(word.RS)-len(suffix1),
		"iéndo", "iendo", "yendo", "ando", "ándo",
		"ár", "ér", "ír", "ar", "er", "ir",
	)
	switch suffix2 {
	case "":

		// Nothing to do
		return false

	case "iéndo", "ándo", "ár", "ér", "ír":

		// In these cases, deletion is followed by removing
		// the acute accent (e.g., haciéndola -> haciendo).

		var suffix2repl string
		switch suffix2 {
		case "":
			return false
		case "iéndo":
			suffix2repl = "iendo"
		case "ándo":
			suffix2repl = "ando"
		case "ár":
			suffix2repl = "ar"
		case "ír":
			suffix2repl = "ir"
		}
		word.ReplaceSuffixRunes(suffix1Runes, []rune(""), true)
		word.ReplaceSuffixRunes(suffix2Runes, []rune(suffix2repl), true)
		return true

	case "ando", "iendo", "ar", "er", "ir":
		word.ReplaceSuffixRunes(suffix1Runes, []rune(""), true)
		return true

	case "yendo":

		// In the case of "yendo", the "yendo" must lie in RV,
		// and be preceded by a "u" somewhere in the word.

		for i := 0; i < len(word.RS)-(len(suffix1)+len(suffix2)); i++ {

			// Note, the unicode code point for "u" is 117.
			if word.RS[i] == 117 {
				word.ReplaceSuffixRunes(suffix1Runes, []rune(""), true)
				return true
			}
		}
	}
	return false
}
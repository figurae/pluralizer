package pluralize

import (
	"strings"
)

type nounForm struct {
	singular string
	plural   string
}

var irreguralNouns = []nounForm{
	{"woman", "women"},
	{"man", "men"},
	{"monarch", "monarchs"},
	{"potato", "potatoes"},
}

func ToPlural(noun string) string {
	for _, irreguralNoun := range irreguralNouns {
		if noun == irreguralNoun.singular {
			return irreguralNoun.plural
		}
	}

	if hasAnySuffix(noun, "s", "x", "z", "j", "ch", "sh") {
		return noun + "es"
	}

	nLen := len(noun)
	if strings.HasSuffix(noun, "y") && nLen > 1 && isConsonant(noun[nLen-2]) {
		return noun[:nLen-1] + "ies"
	}

	return noun + "s"
}

func hasAnySuffix(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}

	return false
}

func isConsonant(char byte) bool {
	const consonants = "bcdfghjklmnpqrstvwxyz"

	return strings.IndexByte(consonants, char) != -1
}

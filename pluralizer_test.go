package pluralizer

import "testing"

var testNouns = []nounForm{
	{"cat", "cats"},
	{"church", "churches"},
	{"monarch", "monarchs"},
	{"butterfly", "butterflies"},
	{"boy", "boys"},
	{"woman", "women"},
	{"potato", "potatoes"},
	{"photo", "photos"},
	{"radio", "radios"},
}

func TestToPlural(t *testing.T) {
	for _, testNoun := range testNouns {
		got := ToPlural(testNoun.singular)
		want := testNoun.plural
		if got != want {
			t.Errorf("expected the plural of %s to be %s but got %s", testNoun.singular, want, got)
		}
	}
}

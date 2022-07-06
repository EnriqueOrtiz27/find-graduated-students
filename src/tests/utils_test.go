package tests

import (
	"findStudent/src/utils"
	"testing"
)

func TestMatchNames(t *testing.T) {
	// should match equal names when transformed to lower case
	got := utils.MatchNames("Alfredo Gonzales", "alfredo", "gonzales")
	if got != true {
		t.Errorf("got %v, wanted %v", got, true)
	}

	// should match people with composite last names
	got_ := utils.MatchNames("Eduardo De Toledo", "Eduardo", "de toledo")
	if got_ != true {
		t.Errorf("got %v, wanted %v", got, true)
	}
}

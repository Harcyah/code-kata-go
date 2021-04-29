package number_adding_one

import (
	"testing"
)

func TestNumberAddingOne(t *testing.T) {
	runNumberAddingOne(8, "9", t)
	runNumberAddingOne(10, "21", t)
	runNumberAddingOne(998, "10109", t)
	runNumberAddingOne(999999, "101010101010", t)
}

func runNumberAddingOne(value int, expected string, t *testing.T) {
	actual := NumberAddingOne(value)
	if actual != expected {
		t.Fatalf("Expecting %d to return %s, but %s was found", value, expected, actual)
	}
}

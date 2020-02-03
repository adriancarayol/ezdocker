package utils

import (
	"strings"
	"testing"
)

func TestOrderString(t *testing.T) {
	expected := "adm"
	input := "mad"

	output := OrderString(input)

	if strings.Compare(expected, output) != 0 {
		t.Fatalf("Fail. Expected: %s, got: %s", expected, output)
	}
}

func TestOrderStringWithSpaces(t *testing.T) {
	expected := " 123adm"
	input := "mad 321"

	output := OrderString(input)

	if strings.Compare(expected, output) != 0 {
		t.Fatalf("Fail. Expected: %s, got: %s", expected, output)
	}
}

func TestOrderStringEmpty(t *testing.T) {
	expected := ""
	input := ""

	output := OrderString(input)

	if strings.Compare(expected, output) != 0 {
		t.Fatalf("Fail. Expected: %s, got: %s", expected, output)
	}
}
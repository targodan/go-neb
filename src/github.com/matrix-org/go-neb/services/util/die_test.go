package util

import "testing"

func TestParseDice(t *testing.T) {
	dice := ParseDice("d6 3d20 9d10")

	if dice.Len() != 1+3+9 {
		t.Errorf("Expected %d dice, but got %d.", 1+3+9, dice.Len())
	}

	expected := "d6, d20, d20, d20, d10, d10, d10, d10, d10, d10, d10, d10, d10"
	result := dice.String()

	if result != expected {
		t.Errorf("Parsing failed. Expected dice:\n%s\nGot:\n%s", expected, result)
	}
}

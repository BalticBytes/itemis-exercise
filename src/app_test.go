package main

import (
	"testing"
)

func TestProcessInput(t *testing.T) {
	input := `glob is I
	prok is V
	pish is X
	tegj is L
	glob glob Silver is 34 Credits
	glob prok Gold is 57800 Credits
	pish pish Iron is 3910 Credits
	how much is pish tegj glob glob ?
	how many Credits is glob prok Silver ?
	how many Credits is glob prok Gold ?
	how many Credits is glob prok Iron ?
	how much wood could a woodchuck chuck if a woodchuck could chuck wood ?`

	expectedOutput := `pish tegj glob glob is 42
	glob prok Silver is 68 Credits
	glob prok Gold is 57800 Credits
	glob prok Iron is 782 Credits
	I have no idea what you are talking about`

	// The same package makes the function visible without exporting it
	output := processInput(input)

	if output != expectedOutput {
		t.Errorf("processInput(%q) = %q, expected %q", input, output, expectedOutput)
	}
}

type TestCase struct {
	romanNumeral  string
	arabicNumeral int
}

func TestRomanNumeral(t *testing.T) {

	fakeIntergalacticMap := map[string]string{
		"I": "I",
		"V": "V",
		"X": "X",
		"L": "L",
		"C": "C",
		"D": "D",
		"M": "M",
	}
	inputs := []TestCase{
		{"I", 1},
		{"I I", 2},
		{"I I I", 3},
		{"I V", 4},
		{"V", 5},
		{"V I", 6},
		{"V I I", 7},
		{"V I I I", 8},
		{"I X", 9},
		{"X", 10},
		{"X I", 11},
		{"X I I", 12},
		{"X I I I", 13},
		{"X I V", 14},
		{"X V", 15},
		{"X V I", 16},
		{"X V I I", 17},
		{"X V I I I", 18},
		{"X I X", 19},
		{"X X", 20},
		{"X I X", 29},
		{"X X X", 30},
		{"X L", 40},
		{"L", 50},
		{"X C", 90},
		{"C", 100},
		{"C D", 400},
		{"D", 500},
		{"C M", 900},
		{"M", 1000},
		{"M M", 2000},
		{"M M M", 3000},
		{"M M M C M X C", 3990},
		{"M M M C M X C I", 3991},
		{"M M M C M X C I I", 3992},
		{"M M M C M X C I I I", 3993},
		{"M M M C M X C I V", 3994},
		{"M M M C M X C V", 3995},
		{"M M M C M X C V I", 3996},
		{"M M M C M X C V I I", 3997},
		{"M M M C M X C V I I I", 3998},
		{"M M M C M X C I X", 3999},
	}

	for _, pair := range inputs {
		numeral := pair.romanNumeral
		expected := pair.arabicNumeral
		_, actual := translate(fakeIntergalacticMap, numeral)
		if expected != actual {
			t.Errorf("translate(%q) = %q, expected %q", numeral, actual, expected)
		}
	}
}

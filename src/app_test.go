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
		t.Errorf("processInput(%s) = '%s'\nexpected %s", input, output, expectedOutput)
	}
}

type TranslateTestCaseInput struct {
	romanNumeral  string
	arabicNumeral int
}

func TestTranslate(t *testing.T) {

	fakeIntergalacticMap := map[string]string{
		"I": "I",
		"V": "V",
		"X": "X",
		"L": "L",
		"C": "C",
		"D": "D",
		"M": "M",
	}
	inputs := []TranslateTestCaseInput{
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
		{"X X I X", 29},
		{"X X X", 30},
		{"X L", 40},
		{"X L I", 41},
		{"X L I I", 42},
		{"X L I X", 49},
		{"L", 50},
		{"L I V", 54},
		{"L I X", 59},
		{"L X I V", 64},
		{"L X I X", 69},
		{"L X X I V", 74},
		{"L X X I X", 79},
		{"L X X X I V", 84},
		{"L X X X I X", 89},
		{"X C", 90},
		{"C", 100},
		{"X C I V", 94},
		{"X C I X", 99},
		{"C I V", 104},
		{"C I X", 109},
		{"C X I V", 114},
		{"C X I X", 119},
		{"C X X I V", 124},
		{"C X X I X", 129},
		{"C X X X I V", 134},
		{"C X X X I X", 139},
		{"C X L I V", 144},
		{"C X L I X", 149},
		{"C L I V", 154},
		{"C L I X", 159},
		{"C L X I V", 164},
		{"C L X I X", 169},
		{"C L X X I V", 174},
		{"C L X X I X", 179},
		{"C L X X X I V", 184},
		{"C L X X X I X", 189},
		{"C X C I V", 194},
		{"C X C I X", 199},
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
			t.Errorf("translate(%q) = %d, expected %d", numeral, actual, expected)
		}
	}
}

type CreditRegexCase struct {
	input   string
	amount  string
	unit    string
	credits string
}

func TestCreditRegexWorks(t *testing.T) {
	inputs := []CreditRegexCase{
		{
			"glob glob Silver is 34 Credits",
			"glob glob",
			"Silver",
			"34",
		},
		// {"glob prok Gold is 57800 Credits", []string{}},
		// {"pish pish Iron is 3910 Credits", []string{}},
	}

	for _, pair := range inputs {
		input := pair.input
		matches := creditsRegex.FindStringSubmatch(input)

		amount := matches[1]
		unit := matches[2]
		credits := matches[3]

		if amount != pair.amount {
			t.Errorf("creditRegex(amount) %s actual %s, expected %s", input, amount, pair.amount)
		}
		if unit != pair.unit {
			t.Errorf("creditRegex(unit) %s actual %s, expected %s", input, unit, pair.unit)
		}
		if credits != pair.credits {
			t.Errorf("creditRegex(credits) %s actual %s, expected %s", input, credits, pair.credits)
		}
	}
}

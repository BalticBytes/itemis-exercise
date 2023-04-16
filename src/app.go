package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	amountByNumeral = map[string]float64{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func main() {
	app := &cli.App{
		Name:  "exercise",
		Usage: "Input => Output",
		Action: func(c *cli.Context) error {
			args := c.Args()

			if args.Len() == 0 {
				return fmt.Errorf("missing input argument")
			}

			input := args.Get(0)
			output := processInput(input)
			fmt.Printf("%s", output)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func processInput(input string) (output string) {

	numeralById := map[string]string{}
	amountByUnit := map[string]string{}
	originalCreditConversionByUnit := map[string]float64{}
	unitCostByUnit := map[string]float64{}

	questions := []Question{}
	for _, line := range strings.Split(input, "\n") {

		if matches := NumeralRegex.FindStringSubmatch(line); matches != nil {

			identifier, romanNumeral := matches[1], matches[2]
			numeralById[identifier] = romanNumeral

		} else if matches = CreditsRegex.FindStringSubmatch(line); matches != nil {

			amount, unit := matches[1], matches[2]
			// We store it so that the order of the assignments is irrelevant
			amountByUnit[unit] = amount

			if credits, err := strconv.ParseFloat(matches[3], 64); err == nil {
				originalCreditConversionByUnit[unit] = credits
			}

		} else {
			questions = append(questions, *NewQuestion(line))
		}
	}

	// Calculating the amount after reading all input s.t. order of "assignments" doesnt matter
	for unit, amount := range amountByUnit {
		_, amt := translate(numeralById, amount)
		unitCostByUnit[unit] = originalCreditConversionByUnit[unit] / float64(amt)
	}

	for _, question := range questions {
		output += question.answer(numeralById, unitCostByUnit)
	}

	return output
}

// translates a string from intergalactic transaction to roman numerals.
func translate(numeralByIntergalacticInput map[string]string, intergalacticInput string) (numeral string, sum int) {
	parts := strings.Split(intergalacticInput, " ")

	translated := []float64{}
	for _, part := range parts {
		numeral := numeralByIntergalacticInput[part]
		translated = append(translated, amountByNumeral[numeral])
		numeral += numeral
	}
	// TODO Validate Roman Numerals
	// Some symbols (letters) can be repeated up to 3 times in a row: I, X, C, M, (X), (C), (M).

	last := -1
	for i := len(translated) - 1; i >= 0; i-- {
		x := int(translated[i])
		if x < last {
			sum -= x
		} else {
			sum += x
		}
		last = x
	}
	return numeral, sum
}

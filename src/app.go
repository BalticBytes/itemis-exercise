package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	numeralPattern = "^[ \t]*([a-zA-Z]+)[ \t]+is[ \t]+([IVXLCDM]+)[ \t]*$"
	creditsPattern = "^[ \t]*([a-zA-Z ]+)[ \t]+([a-zA-Z]+)[ \t]+is[ \t]+([0-9]+)[ \t]+[cC][rR][eE][dD][iI][tT][sS][ \t]*$"
	howMuchPattern = "^[ \t]*how[ \t]+much[ \t]+is[ \t]+([a-zA-Z ]+)[ \t]*\\?$"
	howManyPattern = "^[ \t]*how[ \t]+many[ \t]+[cC][rR][eE][dD][iI][tT][sS][ \t]+is[ \t]+([a-zA-Z ]+)[ \t]+([a-zA-Z]+)[ \t]*\\?$"
)

var (
	numeralRegex = regexp.MustCompile(numeralPattern)
	creditsRegex = regexp.MustCompile(creditsPattern)
	howMuchRegex = regexp.MustCompile(howMuchPattern)
	howManyRegex = regexp.MustCompile(howManyPattern)

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

// There is no "true" enum type so keyword is just a string
type Question struct {
	keyword string // much | many
	amount  string
	unit    *string // optional
}

func (q Question) String() string {
	unit := "N/A"
	if q.unit != nil {
		unit = *q.unit
	}
	return fmt.Sprintf("{How %s is %s %s ?}", q.keyword, q.amount, unit)
}

// returns the sum of an intergalactic
func (q Question) calculate(m map[string]string, n map[string]float64) int {
	_, i := translate(m, q.amount)

	return int(float64(i) * n[*q.unit])
}

// translates a string from intergalactic transaction to roman numerals.
func translate(numeralByIntergalacticInput map[string]string, intergalacticInput string) (string, int) {
	parts := strings.Split(intergalacticInput, " ")
	romanNumerals := ""
	translated := []float64{}
	for _, part := range parts {
		numeral := numeralByIntergalacticInput[part]
		translated = append(translated, amountByNumeral[numeral])
		romanNumerals += numeral
	}
	// TODO Validate Roman Numerals
	// Some symbols (letters) can be repeated up to 3 times in a row: I, X, C, M, (X), (C), (M).
	sum := 0
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
	return romanNumerals, sum
}

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

func processInput(input string) string {

	numeralById := map[string]string{}
	amountByUnit := map[string]string{}
	originalCreditConversionByUnit := map[string]float64{}
	unitCostByUnit := map[string]float64{}

	questions := []Question{}
	var matches []string
	for _, line := range strings.Split(input, "\n") {

		if matches = numeralRegex.FindStringSubmatch(line); matches != nil {
			identifier := matches[1]
			romanNumeral := matches[2]
			numeralById[identifier] = romanNumeral
		} else if matches = creditsRegex.FindStringSubmatch(line); matches != nil {
			amount := matches[1]
			unit := matches[2]

			// We store it so that the order of the assignments is irrelevant
			amountByUnit[unit] = amount

			if credits, err := strconv.ParseFloat(matches[3], 64); err == nil {
				originalCreditConversionByUnit[unit] = credits
			}
		} else if matches = howMuchRegex.FindStringSubmatch(line); matches != nil {
			amount := strings.TrimSpace(matches[1])
			questions = append(questions, Question{"much", amount, nil})
		} else if matches = howManyRegex.FindStringSubmatch(line); matches != nil {
			amount := strings.TrimSpace(matches[1])
			unit := strings.TrimSpace(matches[2])
			questions = append(questions, Question{"many Credits", amount, &unit})
		} else {
			// handle unknown question type
			questions = append(questions, Question{"not a question", "", nil})
		}
	}

	// calculate the amount

	for unit, amount := range amountByUnit {
		_, amt := translate(numeralById, amount)
		unitCostByUnit[unit] = originalCreditConversionByUnit[unit] / float64(amt)
		// fmt.Printf("%12s is %12.3f Credits because %4d x %.0f \n", unit, unitCostByUnit[unit], amt, originalCreditConversionByUnit[unit])
	}

	output := ""
	for _, question := range questions {
		if strings.Contains(question.keyword, "much") {
			// Calculate Roman Numeral
			_, amt := translate(numeralById, question.amount)
			output += fmt.Sprintf("%s is %d\n", question.amount, amt)
		} else if strings.Contains(question.keyword, "many") {
			// Transform
			credits := question.calculate(numeralById, unitCostByUnit)
			output += fmt.Sprintf("%s %s is %d Credits\n", question.amount, *question.unit, credits)
		} else {
			output += "I have no idea what you are talking about"
		}
	}

	return output
}

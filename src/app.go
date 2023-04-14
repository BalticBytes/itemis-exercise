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
	numeralPattern  = "^([a-zA-Z]+) is ([IVXLCDM]+)$"
	creditsPattern  = "^([a-zA-Z ]+) ([a-zA-Z]+) is ([0-9]+) [cC][rR][eE][dD][iI][tT][sS]$"
	questionPattern = "^how (many [cC][rR][eE][dD][iI][tT][sS]|much) is ([a-zA-Z ]+)\\s+([a-zA-Z]*)\\s+\\?$"
)

var (
	numeralRegex  = regexp.MustCompile(numeralPattern)
	creditsRegex  = regexp.MustCompile(creditsPattern)
	questionRegex = regexp.MustCompile(questionPattern)

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

func (q Question) calculate(m map[string]string) int {
	s, i := translate(m, q.amount)
	fmt.Println(q.amount, s, i)
	return i
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
		Name:   "exercise",
		Usage:  "Input => Output",
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	args := c.Args()

	if args.Len() == 0 {
		return fmt.Errorf("missing input argument")
	}

	input := args.Get(0)
	processInput(input)
	// output := processInput(input)
	// fmt.Printf("%s", output)

	return nil
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
		}

		if matches = creditsRegex.FindStringSubmatch(line); matches != nil {
			amount := matches[1]
			unit := matches[2]

			// We store it so that the order of the assignments is irrelevant
			amountByUnit[unit] = amount

			if credits, err := strconv.ParseFloat(matches[3], 64); err == nil {
				originalCreditConversionByUnit[unit] = credits
			}
		}

		if matches = questionRegex.FindStringSubmatch(line); matches != nil {
			keyword := matches[1]
			amount := matches[2]
			unit := matches[3]
			questions = append(questions, Question{
				keyword,
				amount,
				&unit,
			})
		}
	}

	// calculate the amount

	for unit, amount := range amountByUnit {
		xs := strings.Split(amount, " ")
		amt := float64(0)
		// Some symbols (letters) can be repeated up to 3 times in a row: I, X, C, M, (X), (C), (M).
		for i := 0; i < len(xs); i++ {
			x := xs[i]
			y := numeralById[x]
			z := amountByNumeral[y]
			amt += z
		}
		unitCostByUnit[unit] = originalCreditConversionByUnit[unit] / amt
		fmt.Printf("%12s is %8.3f Credits because %3.0f x %.0f \n", unit, originalCreditConversionByUnit[unit]/amt, amt, originalCreditConversionByUnit[unit])
	}

	for _, question := range questions {
		credits := question.calculate(numeralById)
		if strings.Contains(question.keyword, "much") {
			// Calculate Roman Numeral
		} else {
			// Transform
		}
		fmt.Printf("%s %s is %d Credits\n", question.amount, *question.unit, credits)
	}

	return input
}

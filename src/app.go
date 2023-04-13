package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	// "github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)



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

	numeral_pattern := "^([a-zA-Z]+) is ([IVXLCDM]+)$"
	credits_pattern := "^([a-zA-Z ]+) ([a-zA-Z]+) is ([0-9]+) [cC][rR][eE][dD][iI][tT][sS]$"
	question_pattern := "^how (many [cC][rR][eE][dD][iI][tT][sS]|much) is ([a-zA-Z ]+)\\s+([a-zA-Z]*)\\s+\\?$"

	numeral_re := regexp.MustCompile(numeral_pattern)
	credits_re := regexp.MustCompile(credits_pattern)
	question_re := regexp.MustCompile(question_pattern)

	amount_by_numeral := map[string]float64{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	numeral_by_id := map[string]string{}
	amount_by_unit := map[string]string{}
	original_credit_conversion_by_unit := map[string]float64{}
	unitCost_by_unit := map[string]float64{}

	questions := []string{}
	var matches []string
	for _, line := range strings.Split(input, "\n") {

		if matches = numeral_re.FindStringSubmatch(line); matches != nil {
			identifier := matches[1]
			romanNumeral := matches[2]
			numeral_by_id[identifier] = romanNumeral
		}

		if matches = credits_re.FindStringSubmatch(line); matches != nil {
			amount := matches[1]
			unit := matches[2]

			// We store it so that the order of the assignments is irrelevant
			amount_by_unit[unit] = amount

			if credits, err := strconv.ParseFloat(matches[3], 64); err == nil {
				original_credit_conversion_by_unit[unit] = credits
			}
		}

		if matches = question_re.FindStringSubmatch(line); matches != nil {
			keyword := matches[1]
			amount := matches[2]
			unit := matches[3]
			questions = append(questions, line)

			if keyword == "much" {
				fmt.Printf("%s => %s %s\n", line, amount, unit)
			} else {

			}
			fmt.Printf("'%s', '%s', '%s'\n", keyword, amount, unit)
		}
	}

	os.Exit(0)

	// calculate the amount

	for unit, amount := range amount_by_unit {
		xs := strings.Split(amount, " ")
		amt := float64(0)
		// Some symbols (letters) can be repeated up to 3 times in a row: I, X, C, M, (X), (C), (M).
		for i := 0; i < len(xs); i++ {
			x := xs[i]
			y := numeral_by_id[x]
			z := amount_by_numeral[y]
			amt += z
		}
		unitCost_by_unit[unit] = original_credit_conversion_by_unit[unit] / amt
		fmt.Printf("%12s is %8.3f Credits because %3.0f x %.0f \n", unit, original_credit_conversion_by_unit[unit]/amt, amt, original_credit_conversion_by_unit[unit])
	}

	for k, v := range numeral_by_id {
		fmt.Printf("%s: %-8s\n", k, v)
	}

	for k, v := range original_credit_conversion_by_unit {
		fmt.Printf("%s: %-8d\n", k, v)
	}

	return input
}

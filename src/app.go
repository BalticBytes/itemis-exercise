package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

	pattern := []string{
		"^([a-zA-Z]+) is ([IVXLCDM]+)$", // assignment
	}

	id_by_numeral := map[string]string{
		"I": "",
		"V": "",
		"X": "",
		"L": "",
		"C": "",
		"D": "",
		"M": "",
	}

	for _, line := range strings.Split(input, "\n") {
		for i := 0; i < len(pattern); i++ {
			pattern := pattern[i]
			re := regexp.MustCompile(pattern)

			matches := re.FindStringSubmatch(line)

			if matches != nil {
				identifier := matches[1]
				romanNumeral := matches[2]
				// fmt.Printf("%s: %s\n", identifier, romanNumeral)
				id_by_numeral[romanNumeral] = identifier
			}
		}
	}

	for k, v := range id_by_numeral {
		fmt.Printf("%s: %s\n", k, v)
	}

	return input
}

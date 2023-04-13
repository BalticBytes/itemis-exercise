package main

import (
	"fmt"
	"log"
	"os"

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
	output := processInput(input)
	fmt.Printf("%s", output)

	return nil
}

func processInput(input string) string {
	// TODO
	return input
}

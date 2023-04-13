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
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			args := c.Args()

			if args.Len() == 0 {
				return fmt.Errorf("missing input argument")
			}

			input := args.Get(0)
			fmt.Printf("%s", input)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

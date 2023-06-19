package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sethvargo/go-password/password"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Usage = "Random password generation"
	app.Name = "passgen"
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:  "length",
			Value: 16,
			Usage: `Specifies the total length of the generated password. 
	It should be a positive integer.`,
		},
		&cli.IntFlag{
			Name:  "digits",
			Value: 3,
			Usage: `Indicates the minimum number of digits (0-9) that should be included in the generated password. 
	If you set numDigits to 2, the password will have at least 2 digits.`,
		},
		&cli.IntFlag{
			Name:  "symbols",
			Value: 5,
			Usage: `Represents the minimum number of symbols that should be included in the generated password. 
	Symbols are special characters like !@#$%^&*(). 
	If you set numSymbols to 3, the password will have at least 3 symbols.`,
		},
		&cli.BoolFlag{
			Name:  "upper",
			Value: true,
			Usage: `Specifies whether capital letters should be included in the generated password.`,
		},
	}

	app.Action = func(c *cli.Context) error {
		length := c.Int("length")
		digits := c.Int("digits")
		symbols := c.Int("symbols")
		noUpper := !c.Bool("upper")
		allowRepeat := false

		res, err := password.Generate(length, digits, symbols, noUpper, allowRepeat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

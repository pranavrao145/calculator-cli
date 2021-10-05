package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp() // instantiate a new app

func main() {
    // set app properties
    app.Name = "Calculator CLI"
    app.Usage = "A simple calculator CLI written in Golang."
    app.Author = "cypher"
    app.Version = "1.0.0"


    // set up app commands
    app.Commands = []cli.Command{
        {
        	Name:        "add",
        	Aliases:     []string{"a"},
        	Usage:       "add {numbers (separated by space)}",
        	Description: "Adds any amount of numbers together and prints out the result.",
            Action: nil,
        },
        {
        	Name:        "subtract",
        	Aliases:     []string{"sub", "s"},
        	Usage:       "subtract [number] [number]",
        	Description: "Subtracts the second given number from the first and prints out the result.",
            Action: nil,
        },
        {
        	Name:        "multiply",
        	Aliases:     []string{"m", "mult"},
        	Usage:       "mult {numbers (separated by space)}",
        	Description: "Multiplies any amount of numbers together and prints out the result.",
            Action: nil,
        },
        {
        	Name:        "divide",
        	Aliases:     []string{"div", "d"},
        	Usage:       "divide [number] [number]",
        	Description: "Divides the first given number by the second and prints out the result.",
            Action: nil,
        },
        {
        	Name:        "exponent",
        	Aliases:     []string{"exp", "e"},
        	Usage:       "exponent [number] [number]",
        	Description: "Calculates the first number to the exponent of the second and prints out the result.",
            Action: nil,
        },
        {
        	Name:        "logarithm",
        	Aliases:     []string{"log", "l"},
        	Usage:       "logarithm [base] [argument]",
        	Description: "Calculates the logarithm with the base and argument given and prints out the result.",
            Action: nil,
        },
        {
        	Name:        "modulus",
        	Aliases:     []string{"mod", "m"},
        	Usage:       "modulus [number] [number]",
        	Description: "Applies the modulus function the first given number by the second and prints out the result.",
            Action: nil,
        },
    }

    // run app and check for errors
    err := app.Run(os.Args)

    if err != nil {
        log.Fatal(err)
    }
}

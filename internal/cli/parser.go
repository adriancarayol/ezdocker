package cli

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

const minArgs = 2

type Parser struct {}

func New() *Parser {
	return &Parser{}
}

func printHelp() {
	black := color.New(color.FgBlack)
	blackBold := black.Add(color.Bold)
	fmt.Print("Usage: ")
	blackBold.Println("ezd <option> <arguments>")
	fmt.Print("Help: ")
	blackBold.Println("ezd help")
}

// Parse arguments from args
func (p *Parser) ParseOptions() {
	args := os.Args

	if len(args) < minArgs {
		printHelp()
		return
	}

	command := args[1]
	arguments := args[2:]

	parsedArgs := parseParameters(arguments)

	for _, option := range options {
		if strings.Compare(command, option.CommandName) == 0 {
			option.Handler(parsedArgs...)
		}
	}
}

func parseParameters(arguments []string) []string {
	var parsedArgs []string

	for _, arg := range arguments {
		if len(arg) < 2 {
			fmt.Printf("Invalid option: %s\n", arg)
			return parsedArgs
		}

		if strings.HasPrefix(arg, "-") {
			for _, singleArg := range arg[1:] {
				parsedArgs = append(parsedArgs, string(singleArg))
			}
		} else {
			fmt.Printf("Invalid option: %s\n", arg)
			return parsedArgs
		}
	}
	return parsedArgs
}

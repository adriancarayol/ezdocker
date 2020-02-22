package cli

import (
	"fmt"
	"os"
	"strings"
)

const minArgs = 2

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func printHelp() {
	fmt.Print("Usage: ")
	fmt.Println("ezd <option> <arguments>")
	fmt.Print("Help: ")
	fmt.Println("ezd help")
}

// Parse arguments from args
func (p Parser) ParseOptions() {
	args := os.Args

	if len(args) < minArgs {
		printHelp()
		return
	}

	command := args[1]
	arguments := args[2:]

	parsedArgs := p.parseParameters(arguments)

	for _, option := range options {
		if strings.Compare(command, option.CommandName) == 0 {
			option.Handler(parsedArgs...)
		}
	}
}

func (p Parser) parseParameters(arguments []string) []string {
	var parsedArgs []string

	for _, arg := range arguments {
		/*
			Extract complex arguments also,
			for example: -axb will be: -a -x -b
		*/
		if strings.HasPrefix(arg, "-") {
			for _, singleArg := range arg[1:] {
				prefixedArg := "-" + string(singleArg)
				parsedArgs = append(parsedArgs, prefixedArg)
			}
		} else {
			parsedArgs = append(parsedArgs, arg)
		}
	}
	return parsedArgs
}

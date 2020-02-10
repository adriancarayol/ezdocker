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
	msg := blackBold.Sprint("ezd <option> <arguments>")
	fmt.Println(msg)
	fmt.Print("Help: ")
	msg = blackBold.Sprint("ezd help")
	fmt.Println(msg)
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

	parsedArgs, err := parseParameters(arguments)

	if err != nil {
		switch e := err.(type) {
		case *InvalidArgsError:
			fmt.Print(e.Error())
		}
		return
	}

	for _, option := range options {
		if strings.Compare(command, option.CommandName) == 0 {
			option.Handler(parsedArgs...)
		}
	}
}

func parseParameters(arguments []string) ([]string, error) {
	var parsedArgs []string

	for _, arg := range arguments {
		if len(arg) < 2 {
			err := &InvalidArgsError{InvalidArg: arg}
			return nil, err
		}

		if strings.HasPrefix(arg, "-") {
			for _, singleArg := range arg[1:] {
				parsedArgs = append(parsedArgs, string(singleArg))
			}
		} else {
			err := &InvalidArgsError{InvalidArg: arg}
			return nil, err
		}
	}
	return parsedArgs, nil
}

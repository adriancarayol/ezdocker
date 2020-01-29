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

	for _, option := range options {
		if strings.Compare(command, option.CommandName) == 0 {
			option.Handler(args[1:]...)
		}
	}
}

package main

import (
	"github.com/adriancarayol/ezdocker/internal/cli"
)

func main() {
	cli.Init()
	parser := cli.New()
	parser.ParseOptions()
}

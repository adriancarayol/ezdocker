package main

import (
	"github.com/adriancarayol/ezdocker/internal/docker"
)

func main() {
	client := docker.New()
	client.PrintContainers()
}

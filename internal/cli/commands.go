package cli

import (
	"github.com/adriancarayol/ezdocker/internal/docker"
	"log"
)

const (
	list = "ls"
)

var options []Command

func initLsCommand(dockerClient docker.Client) Command {
	client := docker.New(dockerClient)
	printContainersCmd := docker.PrintContainersCommand{Docker: client}
	return Command{CommandName: list, Handler: printContainersCmd.Handle}
}

func ConfigureCommands(dockerClient docker.Client) {
	log.Println("Registering commands...")
	options = append(options, initLsCommand(dockerClient))
	log.Println("Registering commands success.")
}

type CommandHandler func(...string)

type Command struct {
	CommandName string
	Handler     CommandHandler
}

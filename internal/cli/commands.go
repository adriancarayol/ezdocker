package cli

import (
	"github.com/adriancarayol/ezdocker/internal/docker"
	"log"
)

const (
	list = "ls"
)

var options []Command

func initLsCommand() Command {
	dockerClient := docker.New()
	printContainersCmd := docker.PrintContainersCommand{Docker: dockerClient}
	return Command{CommandName: list, Handler: printContainersCmd.Handle}
}

func Init() {
	log.Println("Registering commands...")
	options = append(options, initLsCommand())
	log.Println("Registering commands success.")
}

type CommandHandler func(...string)

type Command struct {
	CommandName string
	Handler     CommandHandler
}

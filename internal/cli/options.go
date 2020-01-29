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
	return Command{CommandName: list, Handler: dockerClient.PrintContainers}
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

package main

import (
	"github.com/adriancarayol/ezdocker/internal/cli"
	"github.com/adriancarayol/ezdocker/internal/docker"
	"github.com/docker/docker/client"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
)

const (
	logFile = "ezdocker.log"
	logFolder = ".ezdocker"
)

// Initialize log files.
func init() {
	userDir, err := homedir.Dir()

	if err != nil {
		log.Printf("Cannot determined the user directory, using stdout for log.")
		return
	}
	fullLogPath := filepath.Join(userDir, logFolder)

	if _, err := os.Stat(fullLogPath); os.IsNotExist(err) {
		errDir := os.Mkdir(fullLogPath, os.ModePerm)
		if errDir != nil {
			log.Printf("Cannot create %s folder, using stdout for log.", fullLogPath)
			return
		}
	}

	fullFileLogPath := filepath.Join(fullLogPath, logFile)

	f, errFile := os.OpenFile(fullFileLogPath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if errFile != nil {
		log.Printf("Cannot create %s log file, using stdout for log. Error: %s", fullFileLogPath, errFile)
		return
	}

	log.SetOutput(f)
}

func initializeDockerClient() docker.Client {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		log.Fatalf("Error building Docker client: %v", err)
	}

	return c
}

func main() {
	c := initializeDockerClient()
	cli.ConfigureCommands(c)
	parser := cli.New()
	parser.ParseOptions()
}

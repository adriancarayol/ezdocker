package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"strings"
	"sync"
)

const (
	all         = "-a"
	stopDefault = ""
)

// Command to stop containers
type StopContainersCommand struct {
	Docker *Docker
}

func printStopContainersHelp() {
	fmt.Println("stop: <option>s")
	fmt.Println("-a Stop all running containers")
}

func (s StopContainersCommand) stopContainers(containerIds ...string) {
	ctx := context.TODO()
	var wg sync.WaitGroup
	wg.Add(len(containerIds))

	for _, containerId := range containerIds {
		go func(containerId string) {
			defer wg.Done()
			fmt.Printf("Stopping container with ID: %s ...\n", containerId)

			err := s.Docker.client.ContainerStop(ctx, containerId, nil)
			if err != nil {
				fmt.Printf("Error stopping container with ID: %s ...\n", containerId)
			}
		}(containerId)
	}
	wg.Wait()
}

func (s StopContainersCommand) stopAllContainers() {
	ctx := context.TODO()
	containers, err := s.Docker.client.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil || len(containers) <= 0 {
		fmt.Println("Not running containers.")
		return
	}

	var containerIds = make([]string, len(containers))

	for i, container := range containers {
		containerIds[i] = container.ID
	}

	s.stopContainers(containerIds...)
}

func (s StopContainersCommand) ExtractOptionsAndParams(opts ...string) ([]string, []string) {
	var params []string
	var options []string

	for _, opt := range opts {
		if strings.HasPrefix(opt, "-") {
			options = append(options, opt)
		} else {
			params = append(params, opt)
		}

	}

	return options, params
}

func (s StopContainersCommand) Handle(opts ...string) {
	options, params := s.ExtractOptionsAndParams(opts...)
	joinedOpts := strings.Join(options, "")

	switch joinedOpts {
	case all:
		s.stopAllContainers()
	case stopDefault:
		s.stopContainers(params...)
	default:
		printStopContainersHelp()
	}

}

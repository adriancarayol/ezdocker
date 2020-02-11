package docker

import (
	"context"
	"fmt"
	"github.com/adriancarayol/ezdocker/internal/utils"
	"github.com/docker/docker/api/types"
	"sort"
	"strings"
)

const (
	all = "a"
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
	for _, containerId := range containerIds {
		fmt.Printf("Stopping container with ID: %s ...\n", containerId)
		err := s.Docker.client.ContainerStop(ctx, containerId, nil)

		if err != nil {
			fmt.Printf("Error stopping container with ID: %s ...\n", containerId)
		}
	}
}


func (s StopContainersCommand) stopAllContainers() {
	ctx := context.TODO()
	containers, err := s.Docker.client.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		fmt.Println("Not running containers")
		return
	}

	var containerIds = make([]string, len(containers))

	for i, container := range containers {
		containerIds[i] = container.ID
	}

	s.stopContainers(containerIds...)
}

func (s StopContainersCommand) Handle(opts ...string) {
	sort.Slice(opts, func(i, j int) bool {
		return opts[i] < opts[j]
	})

	joinedOpts := strings.Join(opts, "")
	switch utils.OrderString(joinedOpts) {
	case all:
		s.stopAllContainers()
	default:
		printStopContainersHelp()
	}

}
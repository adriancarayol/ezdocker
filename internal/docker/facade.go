package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

type Client interface {
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
}

type Docker struct {
	client Client
}

type CommandParameter struct {
	Name string
}

// Create a new Docker to use
func New(client Client) *Docker {
	return &Docker{client: client}
}


//// Stop containers with id in containerIds
//func (d *Docker) StopContainers(containerIds ...string) {
//	ctx := context.TODO()
//	for _, containerId := range containerIds {
//		fmt.Printf("Stopping container with ID: %s ...\n", containerId)
//		err := d.client.ContainerStop(ctx, containerId, nil)
//
//		if err != nil {
//			fmt.Printf("Error stopping container with ID: %s ...\n", containerId)
//		}
//	}
//}
//
//// Stop all running containers
//func (d *Docker) StopAllContainers() {
//	containers := d.ListContainers()
//
//	if containers == nil {
//		fmt.Println("Not running containers")
//		return
//	}
//
//	var containerIds = make([]string, len(containers))
//
//	for i, container := range containers {
//		containerIds[i] = container.ID
//	}
//
//	d.StopContainers(containerIds...)
//}

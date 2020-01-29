package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	dockerClient "github.com/docker/docker/client"
	"github.com/fatih/color"
	"log"
	"strings"
)

type Docker struct {
	client *dockerClient.Client
}

type CommandParameter struct {
	Name string
}

// Create a new Docker to use
func New() *Docker {
	c, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error building Docker client: %v", err)
	}

	return &Docker{client: c}
}

// List running containers
func (d *Docker) ListContainers() []types.Container {
	ctx := context.TODO()
	containers, err := d.client.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		log.Printf("Error listing containers: %v", err)
		return nil
	}

	return containers
}

func printContainer(c types.Container) {
	containerNames := strings.Join(c.Names, ", ")
	fmt.Println(containerNames)

	black := color.New(color.FgBlack)
	blackBold := black.Add(color.Bold)
	blackBold.Printf("  • ID: %s\n", c.ID)
	blackBold.Printf("  • IMAGE: %s\n", c.Image)
	blackBold.Printf("  • STATUS: %s - %s\n", c.State, c.Status)
	blackBold.Printf("  • COMMAND: %s\n", c.Command)
	blackBold.Println("  • PORTS:")

	for _, port := range c.Ports {
		blackBold.Printf("    • IP: %s\n", port.IP)
		blackBold.Printf("    • Public port: %d\n", port.PublicPort)
		blackBold.Printf("    • Private port: %d\n", port.PrivatePort)
		blackBold.Printf("    • Protocol: %s\n", port.Type)
	}
}

// Print running containers
func (d *Docker) PrintContainers(opts ...string) {
	containers := d.ListContainers()

	if containers != nil {
		fmt.Printf("=== Running %d containers ===\n", len(containers))

		for _, c := range containers {
			printContainer(c)
		}
	} else {
		fmt.Println("There's no containers running.")
	}
}

// Stop containers with id in containerIds
func (d *Docker) StopContainers(containerIds ...string) {
	ctx := context.TODO()
	for _, containerId := range containerIds {
		fmt.Printf("Stopping container with ID: %s ...\n", containerId)
		err := d.client.ContainerStop(ctx, containerId, nil)

		if err != nil {
			fmt.Printf("Error stopping container with ID: %s ...\n", containerId)
		}
	}
}

// Stop all running containers
func (d *Docker) StopAllContainers() {
	containers := d.ListContainers()

	if containers == nil {
		fmt.Println("Not running containers")
		return
	}

	var containerIds = make([]string, len(containers))

	for i, container := range containers {
		containerIds[i] = container.ID
	}

	d.StopContainers(containerIds...)
}

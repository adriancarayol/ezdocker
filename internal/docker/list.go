package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/fatih/color"
	"log"
	"strings"
)

const minimal = "-m"

// Command to print containers
type PrintContainersCommand struct {
	Docker *Docker
}

func (p *PrintContainersCommand) listContainers() []types.Container {
	ctx := context.TODO()
	containers, err := p.Docker.client.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		log.Printf("Error listing containers: %v", err)
		return nil
	}

	return containers
}
func printMinimal(c types.Container) {
	containerNames := strings.Join(c.Names, ", ")
	fmt.Println(containerNames)
	black := color.New(color.FgBlack)
	blackBold := black.Add(color.Bold)
	blackBold.Printf("  • ID: %s\n", c.ID)
	blackBold.Printf("  • IMAGE: %s\n", c.Image)
	blackBold.Printf("  • STATUS: %s - %s\n", c.State, c.Status)
}

func printDefaultContainerInfo(c types.Container) {
	black := color.New(color.FgBlack)
	blackBold := black.Add(color.Bold)
	printMinimal(c)

	blackBold.Printf("  • COMMAND: %s\n", c.Command)
	blackBold.Println("  • PORTS:")

	for _, port := range c.Ports {
		blackBold.Printf("    • IP: %s\n", port.IP)
		blackBold.Printf("    • Public port: %d\n", port.PublicPort)
		blackBold.Printf("    • Private port: %d\n", port.PrivatePort)
		blackBold.Printf("    • Protocol: %s\n", port.Type)
	}
}

func (p *PrintContainersCommand) Handle(opts ...string) {
	containers := p.listContainers()

	if containers != nil && len(containers) > 0 {
		fmt.Printf("=== Running %d containers ===\n", len(containers))

		for _, c := range containers {
			if len(opts) <= 0 {
				printDefaultContainerInfo(c)
			}
			for _, opt := range opts {
				switch opt {
				case minimal:
					printMinimal(c)
				default:
					printDefaultContainerInfo(c)
				}
			}
		}
	} else {
		fmt.Println("There's no containers running.")
	}
}

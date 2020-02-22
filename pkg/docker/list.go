package docker

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/adriancarayol/ezdocker/pkg/utils"
	"github.com/docker/docker/api/types"
)

const (
	empty   = ""
	minimal = "-m"
)

// Command to print containers
type PrintContainersCommand struct {
	Docker *Docker
}

func printListHelp() {
	fmt.Println("ls: <option>s")
	fmt.Println("-m Minimal information (id, image, status)")
}

func (p PrintContainersCommand) listContainers() []types.Container {
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

	fmt.Printf("  • ID: %s\n", c.ID)
	fmt.Printf("  • IMAGE: %s\n", c.Image)
	fmt.Printf("  • STATUS: %s - %s\n", c.State, c.Status)
}

func printDefaultContainerInfo(c types.Container) {
	printMinimal(c)

	fmt.Printf("  • COMMAND: %s\n", c.Command)
	fmt.Println("  • PORTS:")

	for _, port := range c.Ports {
		fmt.Printf("    • IP: %s\n", port.IP)
		fmt.Printf("    • Public port: %d\n", port.PublicPort)
		fmt.Printf("    • Private port: %d\n", port.PrivatePort)
		fmt.Printf("    • Protocol: %s\n", port.Type)
	}
}

func (p PrintContainersCommand) ExtractOptionsAndParams(opts ...string) ([]string, []string) {
	sort.Slice(opts, func(i, j int) bool {
		return opts[i] < opts[j]
	})

	return opts, nil
}

func (p PrintContainersCommand) Handle(opts ...string) {
	containers := p.listContainers()

	if containers != nil && len(containers) > 0 {
		options, _ := p.ExtractOptionsAndParams(opts...)

		joinedOpts := strings.Join(options, "")

		switch utils.OrderString(joinedOpts) {
		case empty:
			for _, c := range containers {
				printDefaultContainerInfo(c)
			}
		case utils.OrderString(minimal):
			for _, c := range containers {
				printMinimal(c)
			}
		default:
			printListHelp()
		}

	} else {
		fmt.Println("There's no containers running.")
	}
}

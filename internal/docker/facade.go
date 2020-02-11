package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"time"
)

type Client interface {
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
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
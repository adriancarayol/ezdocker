package mock

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
)

type DockerClient struct{}

func (m DockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return []types.Container{}, nil
}

func (m DockerClient) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return nil
}

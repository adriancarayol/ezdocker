package mock

import (
	"context"
	"github.com/docker/docker/api/types"
)

type DockerClient struct {}

func (m DockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return []types.Container{}, nil
}
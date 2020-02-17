package mock

import (
	"context"
	"errors"
	"time"

	"github.com/docker/docker/api/types"
)

type DockerClient struct{}

type NotEmptyDockerClient struct{}

type ErrorDockerClient struct{}

func (m DockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return []types.Container{}, nil
}

func (m DockerClient) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return nil
}

func (m NotEmptyDockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return []types.Container{{
		ID:      "testingID",
		Names:   nil,
		Image:   "",
		ImageID: "",
		Command: "",
		Created: 0,
		Ports: []types.Port{
			{IP: "192.0.0.0", PrivatePort: 80, PublicPort: 90, Type: "TCP"},
		},
		SizeRw:     0,
		SizeRootFs: 0,
		Labels:     nil,
		State:      "",
		Status:     "",
		HostConfig: struct {
			NetworkMode string `json:",omitempty"`
		}{},
		NetworkSettings: nil,
		Mounts:          nil,
	}}, nil
}

func (m NotEmptyDockerClient) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return nil
}

func (m ErrorDockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return []types.Container{}, errors.New("fake error")
}

func (m ErrorDockerClient) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return errors.New("fake error")
}

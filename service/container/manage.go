package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type ContainerManager struct {
	dockerClient *client.Client
}

func NewContainerManager() (*ContainerManager, error) {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &ContainerManager{dockerClient}, nil
}

func (m *ContainerManager) Create(image string) (string, error) {
	hostConfig := &container.HostConfig{}
	networkingConfig := &network.NetworkingConfig{}

	resp, err := m.dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: image,
		},
		hostConfig,
		networkingConfig,
		nil, "",
	)

	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (m *ContainerManager) Remove(id string) error {
	return m.dockerClient.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{})
}

func (m *ContainerManager) Start(id string) error {
	return m.dockerClient.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
}

func (m *ContainerManager) Stop(id string) error {
	var timeout = 10

	err := m.dockerClient.ContainerStop(context.Background(), id, container.StopOptions{
		Timeout: &timeout,
	})

	return err
}

package circuit

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
)

type dockerEnvironment struct {
	containerID   string
	options       CreateEnvironmentOptions
	dockerOptions DockerEnvironmentOptions
	image         string
	tag           string
	userData      EnvironmentUserData
	isDead        bool
}

func (de *dockerEnvironment) init() error {
	de.isDead = false
	client := de.dockerOptions.Conn
	hostConfig := container.HostConfig{
		Privileged:  false,
		VolumesFrom: []string{de.dockerOptions.DriverInstance},
	}
	hostConfig.Memory = de.dockerOptions.Memory * 1024 * 1024
	config := container.Config{
		Image:     de.dockerOptions.Image,
		Cmd:       []string{de.options.DriverPath},
		OpenStdin: true,
		StdinOnce: false,
	}
	container, err := client.ContainerCreate(context.Background(), &config, &hostConfig, nil, "")
	if err != nil {
		return err
	}
	de.containerID = container.ID
	err = client.ContainerStart(context.Background(), de.containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (de *dockerEnvironment) GetKind() EnvironmentKind {
	return DockerKind
}

func (de *dockerEnvironment) SetUserData(data EnvironmentUserData) error {
	if de.isDead {
		return ErrorDeadEnvironment
	}
	de.userData = data
	return nil
}

func (de *dockerEnvironment) GetUserData() (EnvironmentUserData, error) {
	if de.isDead {
		return nil, ErrorDeadEnvironment
	}
	return de.userData, nil
}

func (de *dockerEnvironment) GetMetadata() EnvironmentMetadata {
	return EnvironmentMetadata{
		"image":     de.dockerOptions.Image,
		"tag":       de.dockerOptions.Tag,
		"container": de.containerID,
	}
}

func (de *dockerEnvironment) Run(request interface{}) (interface{}, error) {
	if de.isDead {
		return nil, ErrorDeadEnvironment
	}

	return nil, nil
}

func (de *dockerEnvironment) Shutdown() error {
	if de.isDead {
		return ErrorDeadEnvironment
	}
	removeOptions := types.ContainerRemoveOptions{
		Force:       true,
		RemoveLinks: true,
	}
	err := de.dockerOptions.Conn.ContainerRemove(context.Background(), de.containerID, removeOptions)
	de.isDead = true
	return err
}

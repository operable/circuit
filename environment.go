package circuit

import (
	"errors"
	"github.com/docker/engine-api/client"
)

type EnvironmentKind int

const (
	NativeKind EnvironmentKind = iota
	DockerKind
)

type EnvironmentMetadata map[string]string
type EnvironmentUserData interface{}

type Environment interface {
	GetKind() EnvironmentKind
	SetUserData(data EnvironmentUserData) error
	GetUserData() (EnvironmentUserData, error)
	GetMetadata() EnvironmentMetadata
	Run(request interface{}) (interface{}, error)
	Shutdown() error
}

type DockerEnvironmentOptions struct {
	Conn           *client.Client
	Image          string
	Tag            string
	DriverInstance string
	Memory         int64
}

type CreateEnvironmentOptions struct {
	Kind          EnvironmentKind
	DriverPath    string
	DockerOptions DockerEnvironmentOptions
}

var ErrorDeadEnvironment = errors.New("Dead environment")

func CreateEnvironment(options CreateEnvironmentOptions) (Environment, error) {
	return nil, nil
}

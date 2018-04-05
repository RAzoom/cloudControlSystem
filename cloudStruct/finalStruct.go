package cloudStruct

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type Service struct {
	ContainerConfig  *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
}

type Server struct {
	Ip          string
	DockerPort  string
	Role        string
	MaxCPU      int
	FreeCPU     int
	MaxMemory   int
	FreeMemory  int
	ServiceMap map[string]Service
}

type Group struct {
	NginxAddr  string
	ConsulAddr string
	ServerMap  map[string]Server
}

type FinalStruct struct {
	MainNginxAddr string
	GroupMap      map[string]Group
}

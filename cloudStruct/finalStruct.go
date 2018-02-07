package cloudStruct

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type Service struct {
	containerConfig  container.Config
	HostConfig       container.HostConfig
	NetworkingConfig network.NetworkingConfig
}

type Server struct {
	ip          string
	dockerPort  string
	serviceList [] Service
}

type Group struct {
	serverList [] Server
}

type FinalStruct struct {
	groupList [] Group
}

func Install(data FinalStruct) {
	for _, group := range data.groupList {
		for _, server := range group.serverList {
			for _, service := range server.serviceList {
				containerConf := service.containerConfig
				hostConf := service.HostConfig
				netConf := service.NetworkingConfig

				containerConf.Hostname = server.ip
			}
		}
	}
}

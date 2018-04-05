package controllers

import (
	"github.com/docker/docker/client"
	"cloudControlSystem/cloudStruct"
	"golang.org/x/net/context"
	"github.com/rs/xid"
	"log"
	"net/http"
	"cloudControlSystem/utils"
	"strings"
	"github.com/docker/docker/api/types"
	"bytes"
)

func nameGenerator(name string) (string) {
	guid := xid.New()
	name = strings.Replace(name, ":", "_", len(name))
	return name + guid.String()
}

func installService(service cloudStruct.Service) (error) {
	// init client
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	//
	pullResult, err := cli.ImagePull(context.Background(), service.ContainerConfig.Image, types.ImagePullOptions{})

	if err != nil {
		log.Panic(err)
	} else {
		println("============== start ==============")

		buf := new(bytes.Buffer)
		buf.ReadFrom(pullResult)
		newStr := buf.String()
		println(newStr)

		println("==============  end  ==============")
	}

	// create new container on server
	contStart, err := cli.ContainerCreate(
		context.Background(),
		service.ContainerConfig,
		service.HostConfig,
		service.NetworkingConfig,
		nameGenerator(service.ContainerConfig.Image))
	if err != nil {
		log.Panic(err)
	}

	err = cli.ContainerStart(
		context.Background(),
		contStart.ID,
		types.ContainerStartOptions{})

	return err

}

func installScheme(data cloudStruct.FinalStruct) (int, error) {
	// read all struct & install containers to cloud
	status := http.StatusOK
	var err error = nil
	for _, group := range data.GroupMap {
		for _, server := range group.ServerMap {
			for _, service := range server.ServiceMap {
				containerConf := service.ContainerConfig
				containerConf.Hostname = server.Ip
				err = installService(service)
				if err != nil {
					log.Panic(err)
					status = http.StatusInternalServerError
				}
			}
		}
	}
	return status, err
}

func InstallScheme() {
	schema := utils.LoadFullSchema()
	installScheme(schema)
}

package controllers

import (
	"net/http"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type Request struct {
	Image   string
	Version string
}

func ColdStartSystem(rw http.ResponseWriter, req *http.Request) (int, io.ReadCloser) {
	var reques Request
	json.NewDecoder(req.Body).Decode(&reques)

	if reques.Image == "" {
		return http.StatusBadRequest, nil
	}

	if reques.Version == "" {
		reques.Version = "latest"
	}
	image := fmt.Sprintf("%s:%s",
		reques.Image,
		reques.Version)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	result, err := cli.ImagePull(context.Background(), image , types.ImagePullOptions{})

	return http.StatusOK, result
}

package task

import (
	"context"
	"io"
	"log"
	"math"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker struct {
	Client *client.Client
	Config Config
}

type DockerResult struct {
	ContainerId string
	Action      string
	Result      string
	Error       error
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()

	log.Printf("Attempting to pull image %s", d.Config.Image)
	reader, err := d.Client.ImagePull(ctx, d.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v", d.Config.Image, err)
		return DockerResult{
			Error: err,
		}
	}
	io.Copy(os.Stdout, reader)

	restartPolicy := container.RestartPolicy{
		Name: container.RestartPolicyMode(d.Config.RestartPolicy),
	}

	resources := container.Resources{
		Memory:   d.Config.Memory,
		NanoCPUs: int64(d.Config.Cpu * math.Pow(10, 9)),
	}

	containerConfig := container.Config{
		Image:        d.Config.Image,
		Env:          d.Config.Env,
		Tty:          false,
		AttachStdin:  d.Config.AttachStdin,
		AttachStdout: d.Config.AttachStdout,
		AttachStderr: d.Config.AttachStderr,
	}

	hostConfig := container.HostConfig{
		RestartPolicy:   restartPolicy,
		Resources:       resources,
		PublishAllPorts: true,
	}

	resp, err := d.Client.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, d.Config.Name)
	if err != nil {
		log.Printf("Error creating container using image %s: %v", d.Config.Image, err)
		return DockerResult{
			Error: err,
		}
	}

	err = d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		log.Printf("Error when starting container %s: %v", resp.ID, err)
		return DockerResult{
			Error: err,
		}
	}

	out, err := d.Client.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Printf("Error getting logs for container %s: %v", resp.ID, err)
		return DockerResult{
			Error: err,
		}
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return DockerResult{
		ContainerId: resp.ID,
		Action:      "start",
		Result:      "success",
		Error:       nil,
	}
}

func (d *Docker) Stop(id string) DockerResult {
	log.Printf("Attempting to stop container %s", id)

	ctx := context.Background()

	err := d.Client.ContainerStop(ctx, id, container.StopOptions{})
	if err != nil {
		log.Printf("Error stopping container %s: %v", id, err)
		return DockerResult{
			Error: err,
		}
	}

	err = d.Client.ContainerRemove(ctx, id, container.RemoveOptions{})
	if err != nil {
		log.Printf("Error removing container %s: %v", id, err)
		return DockerResult{
			Error: err,
		}
	}

	return DockerResult{
		ContainerId: id,
		Action:      "stop",
		Result:      "success",
		Error:       nil,
	}
}

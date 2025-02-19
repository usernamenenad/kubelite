package main

import (
	"time"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/usernamenenad/kubelite/task"
)

func main() {
	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	d := task.Docker{
		Client: c,
		Config: task.Config{
			Name:         "pg-instance",
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Cmd:          []string{},
			Image:        "postgres",
			ExposedPorts: nat.PortSet{
				"5432": struct{}{},
			},
			Env: []string{
				"POSTGRES_PASSWORD=pgpswd",
			},
		},
	}

	result := d.Run()

	time.Sleep(5 * time.Second)

	d.Stop(result.ContainerId)
}

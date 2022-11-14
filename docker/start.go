package docker

import (
    "github.com/docker/docker/api/types"
    "roller/config"
)

func (engine *DockerEngine) StartContainer(name string) {
    _, project := config.GetRollerProject(name)
    if err := engine.c.ContainerStart(*engine.ctx, project.ContainerId, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }
}
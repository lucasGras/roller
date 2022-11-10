package docker

import (
    "github.com/docker/docker/api/types"
    "ops/rollit/roller"
)

func (engine *DockerEngine) StartContainer(name string) {
    _, project := roller.GetRollerProject(name)
    if err := engine.c.ContainerStart(*engine.ctx, project.ContainerId, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }
}
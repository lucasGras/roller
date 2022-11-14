package docker

import "roller/config"

func (engine *DockerEngine) StopContainer(name string) {
    _, project := config.GetRollerProject(name)
    // nil timeout for now
    err := engine.c.ContainerStop(*engine.ctx, project.ContainerId[:12], nil)
    if err != nil {
        panic(err)
    }
}
package docker

import "ops/rollit/roller"

func (engine *DockerEngine) StopContainer(name string) {
    _, project := roller.GetRollerProject(name)
    // nil timeout for now
    err := engine.c.ContainerStop(*engine.ctx, project.ContainerId[:12], nil)
    if err != nil {
        panic(err)
    }
}
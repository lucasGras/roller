package docker

import (
    "fmt"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "roller/config"
)

func (engine *DockerEngine) RolloutContainer(name string) {
    _, oldProject := config.GetRollerProject(name)

    // Start new container
    resp, createErr := engine.c.ContainerCreate(*engine.ctx, &container.Config{
        Image: oldProject.Image,
        }, nil, nil, nil, name + "-roll")
    if createErr != nil {
        panic(createErr)
    }
    startErr := engine.c.ContainerStart(*engine.ctx, resp.ID, types.ContainerStartOptions{})
    if startErr != nil {
        panic(startErr)
    }
    fmt.Println("Start rollout container.")


    // Redirect traffic to new container

    // Stop old container
    stopErr := engine.c.ContainerStop(*engine.ctx, oldProject.ContainerId[:12], nil)
    if stopErr != nil {
        panic(stopErr)
    }
    fmt.Println("Stop old container.")
    // Delete old container
    delErr := engine.c.ContainerRemove(*engine.ctx, oldProject.ContainerId, types.ContainerRemoveOptions{})
    if delErr != nil {
        panic(delErr)
    }
    // Rename new container
    renameErr := engine.c.ContainerRename(*engine.ctx, resp.ID, name)
    if renameErr != nil {
        panic(renameErr)
    }
    fmt.Println("Clean-up containers.")

    // Update roller project
    err := config.UpdateRollerProject(name, resp.ID, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println("Container roll-out done.")
}
package docker

import "roller/config"

const (
    PRUNE_MODE_ALL = iota
    PRUNE_MODE_SINGLE = iota
)

// Prune running project, remove containers and clear roller cfg
func (engine *DockerEngine) Prune(mode int, name *string) {
    if mode == PRUNE_MODE_ALL && name != nil {
        panic("invalid arguments")
    }
    if mode == PRUNE_MODE_ALL {
        _, projects := config.GetRollerProjects()
        for _, p := range projects {
            if err := engine.c.ContainerStop(*engine.ctx, p.ContainerId[:12], nil); err != nil {
                panic(err)
                // Handle this
            }
            if err := config.DeleteRollerProject(p.Name); err != nil {
                panic(err)
            }
        }
        // Delete all is not safe in case a container fails stopping
    } else if mode == PRUNE_MODE_SINGLE && name != nil {
        if err := config.DeleteRollerProject(*name); err != nil {
            panic(err)
        }
    } else {
        panic("unknown prune behaviour")
    }
}
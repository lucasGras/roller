package docker

import (
    ptable "github.com/jedib0t/go-pretty/v6/table"
    "roller/config"
    "os"
)

// Status displays all roller projects status
//  - project name
//  - port mapping
//  - image
//  - id
//  - roller status (UP | STOPED | EXPOSED)
//  - if exposed we want some more information about reverse proxy and dns
//  With a stats more we can see some info about mem and cpu usage
func (engine *DockerEngine) Status() {
    _, projects := config.GetRollerProjects()
    table := ptable.NewWriter()
    table.SetOutputMirror(os.Stdout)
    table.AppendHeader(ptable.Row{"#", "name", "image", "ports", "status"})

    for _, p := range projects {
        _, err := engine.c.ContainerStatsOneShot(*engine.ctx, p.ContainerId)
        if err != nil {
            panic(err)
        }
        table.AppendRow(ptable.Row{ p.ContainerId[:12], p.Name, p.Image, "3000", "UP" })
        table.AppendSeparator()
    }
    table.SetStyle(ptable.StyleColoredBright)
    table.Render()
}
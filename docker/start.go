package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"net"
	"roller/config"
	"strconv"
	"strings"
)

type StartContainerOpts struct {
	TcpPort  nat.Port
	HostPort string
	HostIP   string
}

func ParseStartContainerOpts(opts map[string]string) StartContainerOpts {
	startContainerOpts := StartContainerOpts{}
	if _, ok := opts["ports"]; ok {
		// port mapping similar to docker
		// HostPort:TcpPort
		// https://docs.docker.com/config/containers/container-networking/
		ports := strings.Split(opts["ports"], ":")
		if len(ports) != 2 {
			panic("malformed ports ")
		}
		if _, err := strconv.ParseUint(ports[0], 10, 64); err != nil {
			panic("malformed host port")
		}
		if _, err := strconv.ParseUint(ports[1], 10, 64); err != nil {
			panic("malformed target port")
		}
		startContainerOpts.HostPort = ports[0]
		startContainerOpts.TcpPort = nat.Port(ports[1] + "/tcp")
	} else {
		panic("ports flag is required for port mapping")
	}
	if _, ok := opts["ip"]; ok {
		if net.ParseIP(opts["ip"]) == nil {
			panic("malformed ip address")
		}
		startContainerOpts.HostIP = opts["ip"]
	} else {
		startContainerOpts.HostIP = "0.0.0.0"
	}
	return startContainerOpts
}

func (engine *DockerEngine) StartContainer(name string, opts StartContainerOpts) {
	var containerId string
	getErr, project := config.GetRollerProject(name)
	if getErr != nil {
		resp, err := engine.c.ContainerCreate(*engine.ctx, &container.Config{
			Image: project.Image,
		}, &container.HostConfig{
			PortBindings: nat.PortMap{
				opts.TcpPort: []nat.PortBinding{
					{HostIP: opts.HostIP, HostPort: opts.HostPort},
				},
			},
		}, nil, nil, name)

		if err != nil {
			panic(err)
		}
		fmt.Println("Create new container.")
		containerId = resp.ID
	} else {
		containerId = project.ContainerId
	}

	fmt.Println(containerId)
	if err := engine.c.ContainerStart(*engine.ctx, project.ContainerId, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	/*
		if rollerErr != nil {
			// Painful error
			// Rescue remanent container
			engine.c.ContainerKill(*engine.ctx, containerId, "SIGABRT")

			fmt.Errorf("Could not finalize container startup. Abort container creation.", rollerErr)
			return
		}
	*/
	config.UpdateRollerProject(name, containerId, nil)
}

package docker

import (
    "context"
    "github.com/docker/docker/client"
)

type DockerEngine struct {
    c *client.Client
    ctx *context.Context
}

func InitDockerEngine(client *client.Client, ctx *context.Context) *DockerEngine {
    return &DockerEngine{
        c: client,
        ctx: ctx,
    }
}

// ➜  rollit ./rollit create roger snaipeberry/bpc_images:backend-test
// ➜  rollit echo '{"projects": []}'  > roller
// caddy reverse-proxy --to :3000
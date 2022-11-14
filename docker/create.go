package docker

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/go-connections/nat"
    "io"
    "os"
    "roller/config"
)

func (engine *DockerEngine) CreateContainer(name string, image string) {
    authConfig := types.AuthConfig{
        Username: "snaipeberry",
        Password: "",
    }
    encodedJSON, err := json.Marshal(authConfig)
    if err != nil {
        panic(err)
    }
    authStr := base64.URLEncoding.EncodeToString(encodedJSON)

    out, err := engine.c.ImagePull(*engine.ctx, image, types.ImagePullOptions{RegistryAuth: authStr})
    if err != nil {
        panic(err)
    }
    defer out.Close()
    io.Copy(os.Stdout, out)
    // If the container already exists with the same name, don't create it
    var containerId string
    getError, project := config.GetRollerProject(name)
    if getError != nil {
        resp, err := engine.c.ContainerCreate(*engine.ctx, &container.Config{
            Image: image,
            }, &container.HostConfig{
            PortBindings: nat.PortMap{
                "3000/tcp": []nat.PortBinding{
                    {HostIP: "0.0.0.0", HostPort: "3000"},
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

    rollerErr := config.CreateRollerProject(name, containerId, image)
    if rollerErr != nil {
        // Painful error
        // Rescue remanent container
        engine.c.ContainerKill(*engine.ctx, containerId, "SIGABRT")

        fmt.Errorf("Could not finalize container startup. Abort container creation.", rollerErr)
        return
    }
}

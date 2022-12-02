package docker

import (
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"io"
	"os"
	"roller/config"
)

type CreateContainerOpts struct {
	User     string
	Password string
}

func ParseCreateContainerOpts(opts map[string]string) CreateContainerOpts {
	return CreateContainerOpts{
		User:     opts["user"],
		Password: opts["password"],
	}
}

func createImagePullOptions(opt CreateContainerOpts) types.ImagePullOptions {
	if opt.User != "" && opt.Password != "" {
		authConfig := types.AuthConfig{
			Username: opt.User,
			Password: opt.Password,
		}
		encodedJSON, err := json.Marshal(authConfig)
		if err != nil {
			panic(err)
		}
		authStr := base64.URLEncoding.EncodeToString(encodedJSON)
		return types.ImagePullOptions{
			RegistryAuth: authStr,
		}
	} else if opt.User != "" || opt.Password != "" {
		panic("user and password are needed to login into dockerhub")
	}
	return types.ImagePullOptions{}
}

func (engine *DockerEngine) CreateContainer(name string, image string, opt CreateContainerOpts) {
	imagePullOptions := createImagePullOptions(opt)
	out, err := engine.c.ImagePull(*engine.ctx, image, imagePullOptions)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	cfgErr := config.CreateRollerProject(name, "", image)
	if cfgErr != nil {
		panic(err)
	}
}

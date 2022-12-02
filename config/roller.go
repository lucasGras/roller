package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Roller struct {
	Projects []RollerProject `json:"projects"`
}

type RollerProject struct {
	ContainerId string `json:"containerId"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	// status byte
}

// /.rollit/projects.json
// Store basic key values file
// [rollit_container_name:container_id]
func GetRollerProject(name string) (error, *RollerProject) {
	data := readFile("/.rollit/roller")
	return findProjectInCfgFile(data, name)
}

func GetRollerProjects() (error, []RollerProject) {
	data := readFile("/.rollit/roller")
	return nil, data.Projects
}

func CreateRollerProject(name string, containerId string, image string) error {
	data := readFile("/.rollit/roller")
	// Check if project does not already exists
	newData := append(data.Projects, RollerProject{
		Name:        name,
		ContainerId: containerId,
		Image:       image,
	})
	data.Projects = newData
	writeFile("/.rollit/roller", data)
	return nil
}

func UpdateRollerProject(name string, containerId string, image *string) error {
	data := readFile("/.rollit/roller")
	_, newData := updateProjectInCfgFile(&data, name, containerId, image)
	writeFile("/.rollit/roller", *newData)
	return nil
}

func DeleteRollerProject(name string) error {
	data := readFile("/.rollit/roller")
	err, newData := deleteProjectInCfgFile(&data, name)
	if err != nil {
		panic(err)
	}
	writeFile("/.rollit/roller", *newData)
	return nil
}

func DeleteRollerProjects() error {
	writeFile("/.rollit/roller", Roller{Projects: []RollerProject{}})
	return nil
}

// ---- core

func writeFile(filePath string, data Roller) {
	home, _ := os.UserHomeDir()
	file, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	writeErr := ioutil.WriteFile(home+filePath, file, 0644)
	if writeErr != nil {
		panic(writeErr)
	}
}

func readFile(filePath string) Roller {
	home, _ := os.UserHomeDir()
	file, err := ioutil.ReadFile(home + filePath)
	if err != nil {
		panic(err)
	}
	data := Roller{}
	jsonErr := json.Unmarshal(file, &data)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return data
}

func findProjectInCfgFile(data Roller, name string) (error, *RollerProject) {
	for _, project := range data.Projects {
		if project.Name == name {
			return nil, &project
		}
	}
	return fmt.Errorf("trying to find a non existing project"), nil
}

// Use array idx instead of value to update reference
func updateProjectInCfgFile(data *Roller, name string, containerId string, image *string) (error, *Roller) {
	for i, _ := range data.Projects {
		if data.Projects[i].Name == name {
			data.Projects[i].ContainerId = containerId
			if image != nil {
				data.Projects[i].Image = *image
			}
			return nil, data
		}
	}

	return fmt.Errorf("trying to find a non existing project"), nil
}

func deleteProjectInCfgFile(data *Roller, name string) (error, *Roller) {
	var idx int = -1

	for i, _ := range data.Projects {
		if data.Projects[i].Name == name {
			idx = i
			break
		}
	}

	if idx < 0 {
		return fmt.Errorf("trying to find a non existing project"), nil
	}
	data.Projects = append(data.Projects[:idx], data.Projects[idx+1:]...)
	return nil, data
}

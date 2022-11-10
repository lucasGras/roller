package docker

const (
    PRUNE_MODE_ALL = iota
    PRUNE_MODE_SINGLE = iota
)

// Prune running project, remove containers and clear roller cfg
func (engine *DockerEngine) Prune(mode int) {

}
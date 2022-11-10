package main

import (
    "context"
    "github.com/caddyserver/caddy/v2"
    "ops/rollit/docker"
    proxy "ops/rollit/reverse-proxy"
    "os"

    "github.com/docker/docker/client"
    "github.com/teris-io/cli"
)

func main() {
    ctx := context.Background()
    dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer dockerClient.Close()
    dockerEngine := docker.InitDockerEngine(dockerClient, &ctx)

    createCmd := cli.NewCommand("create", "create a new repository in rollit").
        WithArg(cli.NewArg("name", "name of your rollit project")).
        WithArg(cli.NewArg("image", "the docker image to use")).
        WithAction(func(args []string, options map[string]string) int {
            var name = args[0]
            var image = args[1]

            dockerEngine.CreateContainer(name, image)
            return 0
        },
    )
    startCmd := cli.NewCommand("start", "start a rollit project").
        WithArg(cli.NewArg("name", "name of your rollit project")).
        WithAction(func(args []string, options map[string]string) int {
            var name = args[0]

            dockerEngine.StartContainer(name)
            return 0
        },
    )
    stopCmd := cli.NewCommand("stop", "stop a rollit project").
        WithArg(cli.NewArg("name", "name of your rollit project")).
        WithAction(func(args []string, options map[string]string) int {
            var name = args[0]

            dockerEngine.StopContainer(name)
            return 0
        },
    )
    rollCmd := cli.NewCommand("roll", "perform project roll-out").
        WithArg(cli.NewArg("name", "name of your rollit project")).
        WithAction(func(args []string, options map[string]string) int {
            var name = args[0]

            dockerEngine.RolloutContainer(name)
            return 0
        },
    )
    pruneCmd := cli.NewCommand("prune", "clean rollit project(s)").
        WithOption(cli.NewOption("all", "clean all rollit projects").WithChar('a').WithType(cli.TypeBool)).
        WithAction(func(args []string, options map[string]string) int {
            var mod = docker.PRUNE_MODE_SINGLE
            if _, ok := options["all"]; ok {
                mod = docker.PRUNE_MODE_ALL
            }
            dockerEngine.Prune(mod)
            return 0
        },
    )
    statusCmd := cli.NewCommand("status", "get information about your projetcs").
        WithAction(func(args []string, options map[string]string) int {
            dockerEngine.Status()
            return 0
        },
    )
    exposeCmd := cli.NewCommand("expose", "expose rollit project").
        WithAction(func(args []string, options map[string]string) int {
            proxy.RunProxy(&caddy.Config{
                AppsRaw: caddy.ModuleMap{
                    "http": proxy.ProjectProxyToModuleMap(proxy.ProxyHttpModule{
                        Servers: []proxy.ProxyHttpModuleServer{
                            {Listen: []string{":3000"}},
                        },
                    }),
                },
            })
            return 0
        },
    )
    app := cli.New("rollit is a super-simple and straigh-forward way to deploy your small projects to the web").
        WithCommand(createCmd).
        WithCommand(startCmd).
        WithCommand(stopCmd).
        WithCommand(rollCmd).
        WithCommand(statusCmd).
        WithCommand(pruneCmd).
        WithCommand(exposeCmd)
    os.Exit(app.Run(os.Args, os.Stdout))
}
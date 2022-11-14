package main

import (
    "context"
    "encoding/json"
    "github.com/caddyserver/caddy/v2"
    "github.com/caddyserver/caddy/v2/modules/caddyhttp"
    _ "github.com/caddyserver/caddy/v2/modules/caddyhttp/standard"
    "os"
    "roller/docker"
    proxy "roller/reverse-proxy"

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
        WithArg(cli.NewArg("name", "name the project to delete").AsOptional()).
        WithOption(cli.NewOption("all", "clean all rollit projects").WithType(cli.TypeBool).WithChar('a')).
        WithAction(func(args []string, options map[string]string) int {
            var mod = docker.PRUNE_MODE_SINGLE

            if _, ok := options["all"]; ok {
                mod = docker.PRUNE_MODE_ALL
                dockerEngine.Prune(mod, nil)
                return 0
            }
            dockerEngine.Prune(mod, &args[0])
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
        WithArg(cli.NewArg("name", "name of your rollit project")).
        WithAction(func(args []string, options map[string]string) int {
            // var name = args[0]

            proxy.RunProxy(&caddy.Config{
                AppsRaw: caddy.ModuleMap{
                    "http": proxy.ProjectProxyToModuleMap(caddyhttp.App{
                        Servers: map[string]*caddyhttp.Server{
                            "roger": &caddyhttp.Server{
                                Listen: []string{":80"},
                                Routes: []caddyhttp.Route{
                                    { HandlersRaw: []json.RawMessage{
                                        []byte(`{
                                                "handler": "reverse_proxy",
                                                "upstreams": [{"dial": ":3000"}]
                                                }`),
                                    }},
                                },
                            },
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
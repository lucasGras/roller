package proxy

import(
    "fmt"
    "github.com/caddyserver/caddy/v2"
)

func RunProxy(cfg *caddy.Config) {
    fmt.Print(*cfg)
    err := caddy.Run(cfg)
    if err != nil {
        panic(err)
    }
}

func StopProxy() {
    err := caddy.Stop()
    if err != nil {
        panic(err)
    }
}
package proxy

import(
    "github.com/caddyserver/caddy/v2"
    web_server "roller/web-server"
)

// Solution for background processing
// - run proxy as a sub process and pid handling (go run proxy.go)
// - handle caddy as a running posix service and just run / load the config here
// - caddy start (sysctl / ...) and use rest api ==> OK
func RunProxy(cfg *caddy.Config) {
    web_server.CaddyLoad(*cfg)

}

func StopProxy() {
    err := caddy.Stop()
    if err != nil {
        panic(err)
    }
}
package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"roller/utils"
)

// Solution for background processing
// - run proxy as a sub process and pid handling (go run proxy.go)
// - handle caddy as a running posix service and just run / load the config here
// - caddy start (sysctl / ...) and use rest api ==> OK
func RunProxy(cfg *caddy.Config) {
	data, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	req := utils.NewRequestWithForm("POST", "load", data)
	response, reqErr := utils.DoRequest(req)
	if reqErr != 200 {
		panic(reqErr)
	}
	fmt.Print(response)
}

func StopProxy() {
	err := caddy.Stop()
	if err != nil {
		panic(err)
	}
}

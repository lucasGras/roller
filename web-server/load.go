package web_server

import (
    "encoding/json"
    "fmt"
    "github.com/caddyserver/caddy/v2"
    "roller/utils"
)

func CaddyLoad(cfg caddy.Config) {
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
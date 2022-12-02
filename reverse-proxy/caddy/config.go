package proxy

import (
	"encoding/json"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func ProjectProxyToModuleMap(projectProxy caddyhttp.App) []byte {
	data, _ := json.Marshal(projectProxy)
	return data
}

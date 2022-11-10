package proxy

import (
    "encoding/json"
    "io/ioutil"
    "os"
)
type ProxyHttpModuleServer struct {
    Listen []string `json:"listen"`
}

type ProxyHttpModule struct {
    Servers []ProxyHttpModuleServer `json:"servers"`
}

func ProjectProxyToModuleMap(projectProxy ProxyHttpModule) []byte {
    data, _ := json.Marshal(projectProxy)
    return data
}

func readConfigFile() {
    home, _ := os.UserHomeDir()
    file, err := ioutil.ReadFile(home + "/.rollit/proxy.json")
    if err != nil {
        panic(err)
    }
    data := ProxyHttpModule{}
    json.Unmarshal(file, &data)
}
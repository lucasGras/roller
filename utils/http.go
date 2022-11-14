package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

func NewRequest(method string, path string) *http.Request {
    req, err := http.NewRequest(method, "http://localhost:2019/"+path, nil)

    if err != nil {
        panic("can not create http request")
    }
    return req
}

func NewRequestWithForm(method string, path string, form []byte) *http.Request {
    req, err := http.NewRequest(method, "http://localhost:2019/"+path, strings.NewReader(string(form)))
    // req.MultipartForm.Value = form
    req.Header.Add("Content-Type", "application/json")
    if err != nil {
        panic("can not create http request")
    }
    return req
}

func DoRequest(req *http.Request) ([]byte, int) {
    c := http.Client{Timeout: time.Duration(3) * time.Second}
    resp, err := c.Do(req)
    if err != nil {
        fmt.Printf("Error %s", err)
        return nil, resp.StatusCode
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Printf("Body : %s", body)
    return body, resp.StatusCode
}

func DoUnmarshall(data []byte, v interface{}) {
    err := json.Unmarshal(data, &v)
    if err != nil {
        panic("Error unmarshalling")
    }
}
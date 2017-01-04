package tlgrmBotApi

import (
//    "encoding/json"; 
    "net/http"; 
    "fmt"; 
    "bytes"
)

var token string

const urlTemplate string = "https://api.telegram.org/bot%s/%s"

func SetToken(extToken string) {
    token = extToken
}

func Call(method string, param []byte) (interface{}, bool) {
    url := fmt.Sprintf(urlTemplate, token, method);
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(param));
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)

    return nil, true
}

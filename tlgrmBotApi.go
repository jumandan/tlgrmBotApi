package tlgrmBotApi

import (
    "encoding/json"; 
    "net/http"
    "fmt"
    "bytes"
    //"io/ioutil"
)

var token string

const urlTemplate string = "https://api.telegram.org/bot%s/%s"

func SetToken(extToken string) {
    token = extToken
}

func Call(method string, param interface{}, result interface{}) {
    url := fmt.Sprintf(urlTemplate, token, method);
    paramStr, err := json.Marshal(param)
    if err != nil {
        panic(err)
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramStr));
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

    if resp.StatusCode != 200 {
        panic("Server error. Status is " + resp.Status)
    }

    decoder := json.NewDecoder(resp.Body)
    decoder.UseNumber()
    if err := decoder.Decode(result); err != nil {
        panic(err)
    }
}

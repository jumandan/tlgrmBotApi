package tlgrmBotApi

import (
    "encoding/json"; 
    "net/http"
    "fmt"
    "bytes"
)

/**
 * private variable
 */
var token string

/**
 * private constant
 */
const urlTemplate string = "https://api.telegram.org/bot%s/%s"

/**
 * Set Bot token
 */
func SetToken(extToken string) {
    token = extToken
}

/**
 * Call API
 */
func Call(method string, param interface{}, result interface{}) {
    url := fmt.Sprintf(urlTemplate, token, method);
    paramStr, err := json.Marshal(param)
    if err != nil {
        panic(err)
    }
    fmt.Printf("params is %s\n", paramStr)
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

package tlgrmBotApi

import (
    "encoding/json"; 
    "net/http"
    "fmt"
    "bytes"
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

type GetUpdatesRecord struct {
    Ok     bool           `json:"ok"`
    Result []updateRecord `json:"result"`
}

type GetMeRecord struct {
    Ok     bool       `json:"ok"`
    Result userRecord `json:"result"`
}

type userRecord struct {
    Id         uint64 `json:"id"`
    Username   string `json:"username"`
    First_name string `json:"first_name"`
    Last_name  string `json:last_name`
}

type messageRecord struct {
    Message_id int64      `json:"message_id"`
    From       userRecord `json:"from"`
    Date       int        `json:"date"`
    Text       string     `json:"text"`
    Chat       chatRecord `json:"chat"`
}

type chatRecord struct {
    Id   int64  `json:"id"`
    Type string `json:"type"`
}

type updateRecord struct {
    Update_id uint64        `json:"id"`
    Message   messageRecord `json:"message"`
}



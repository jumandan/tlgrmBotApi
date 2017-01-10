package tlgrmBotApi

type User struct {
    Id        uint64 `json:"id"`
    UserName  string `json:"username"`
    FirstName string `json:"first_name"`
    LastName  string `json:last_name`
}

type Chat struct {
    Id   int64  `json:"id"`
    Type string `json:"type"`
}

type PhotoSize struct {
    FileId   string `json:"file_id"`
    Width    uint   `json:"width"`
    Height   uint   `json:"height"`
    FileSize uint   `json:"file_size"`
}

type Sticker struct {
    FileId   string    `json:"file_id"`
    Width    uint      `json:"width"`
    Height   uint      `json:"height"`
    Tumb     PhotoSize `json:"tumb"`
    Emoji    string    `json:"emoji"`
    FileSize uint      `json:"file_size"`
}

type Message struct {
    Id      int64   `json:"message_id"`
    From    User    `json:"from"`
    Date    int     `json:"date"`
    Text    string  `json:"text"`
    Chat    Chat    `json:"chat"`
    Sticker Sticker `json:"sticker"`
}

type Update struct {
    UpdateId uint64 `json:"update_id"`
    Message Message `json:"message"`
}

type GetUpdates struct {
    Ok     bool     `json:"ok"`
    Result []Update `json:"result"`
}

type GetMe struct {
    Ok     bool `json:"ok"`
    Result User `json:"result"`
}

type SendMessage struct {
    Ok     bool    `json:"ok"`
    Result Message `json:"result"`
}


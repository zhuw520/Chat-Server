package protocol

import "chat-server/model"

type InitMessage struct {
    Type        string              `json:"type"`
    Messages    []model.ChatMessage `json:"messages"`
    UserNumber  int                 `json:"userNumber"`
    OnlineCount int                 `json:"onlineCount"`
}

type ChatUpdate struct {
    Type        string            `json:"type"`
    Message     model.ChatMessage `json:"message"`
    OnlineCount int               `json:"onlineCount"`
}

type ErrorMessage struct {
    Type    string `json:"type"`
    Message string `json:"message"`
}
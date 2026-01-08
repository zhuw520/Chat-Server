package model

type ChatMessage struct {
    Type       string `json:"type"`
    UserID     string `json:"userId"`
    UserNumber int    `json:"userNumber"`
    Message    string `json:"message"`
    Timestamp  string `json:"timestamp"`
    IP         string `json:"ip"`
}
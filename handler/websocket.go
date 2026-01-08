package handler

import (
    "encoding/json"
    "net/http"
    "chat-server/server"
    "chat-server/protocol"
    "github.com/gorilla/websocket"
)

type WebSocketHandler struct {
    server *server.ChatServer
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func (h *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    clientIP := r.RemoteAddr
    client := &server.Client{
        Conn: conn,
        IP:   clientIP,
    }
    h.server.clientsMutex.Lock()
    h.server.clients[client] = true
    h.server.clientsMutex.Unlock()
    go h.handleClient(client)
}

func (h *WebSocketHandler) handleClient(client *server.Client) {
    defer func() {
        h.server.clientsMutex.Lock()
        delete(h.server.clients, client)
        h.server.clientsMutex.Unlock()
        client.Conn.Close()
    }()
    for {
        _, msgBytes, err := client.Conn.ReadMessage()
        if err != nil {
            break
        }
        var msgData map[string]interface{}
        if err := json.Unmarshal(msgBytes, &msgData); err != nil {
            continue
        }
    }
}
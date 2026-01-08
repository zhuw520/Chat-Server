// Unescaped
package server

import (
	"encoding/json"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
	"chat-server/storage"
	"chat-server/model"
)

type ChatServer struct {
	clients      map[*Client]bool
	clientsMutex sync.RWMutex
	messageStore *storage.MemoryStore
}

type Client struct {
	Conn       *websocket.Conn
	UserID     string
	UserNumber int
	IP         string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients:      make(map[*Client]bool),
		messageStore: storage.NewMemoryStore(),
	}
}

func (s *ChatServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		Conn: conn,
		IP:   r.RemoteAddr,
	}
	go s.handleClient(client)
}

func (s *ChatServer) handleClient(client *Client) {
	defer func() {
		s.clientsMutex.Lock()
		delete(s.clients, client)
		s.clientsMutex.Unlock()
		client.Conn.Close()
	}()
	s.clientsMutex.Lock()
	s.clients[client] = true
	s.clientsMutex.Unlock()
	for {
		_, msgBytes, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
		var msgData map[string]interface{}
		if err := json.Unmarshal(msgBytes, &msgData); err != nil {
			continue
		}
		msgType, _ := msgData["type"].(string)
		if msgType == "join" {
			userID, _ := msgData["userId"].(string)
			client.UserID = userID
			client.UserNumber = 1
			s.sendInitData(client)
			joinMsg := model.ChatMessage{
				Type:       "join",
				UserID:     userID,
				UserNumber: client.UserNumber,
				Message:    "joined",
				IP:         client.IP,
			}
			s.broadcastMessage(joinMsg)
		} else if msgType == "message" && client.UserID != "" {
			message, _ := msgData["message"].(string)
			timestamp, _ := msgData["timestamp"].(string)
			chatMsg := model.ChatMessage{
				Type:       "user",
				UserID:     client.UserID,
				UserNumber: client.UserNumber,
				Message:    message,
				Timestamp:  timestamp,
				IP:         client.IP,
			}
			s.messageStore.Add(chatMsg)
			s.broadcastMessage(chatMsg)
		}
	}
}

func (s *ChatServer) sendInitData(client *Client) {
	initMsg := map[string]interface{}{
		"type":        "init",
		"messages":    s.messageStore.GetAll(),
		"userNumber":  client.UserNumber,
		"onlineCount": len(s.clients),
	}
	data, _ := json.Marshal(initMsg)
	client.Conn.WriteMessage(websocket.TextMessage, data)
}

func (s *ChatServer) broadcastMessage(msg model.ChatMessage) {
	update := map[string]interface{}{
		"type":    "chat_update",
		"message": msg,
	}
	data, _ := json.Marshal(update)
	s.clientsMutex.RLock()
	defer s.clientsMutex.RUnlock()
	for client := range s.clients {
		client.Conn.WriteMessage(websocket.TextMessage, data)
	}
}
package server

import (
    "sync"
    "time"
    "chat-server/model"
)

type ConnectionManager struct {
    onlineUsers   map[string]*model.UserInfo
    onlineMutex   sync.RWMutex
    userNumberMap map[string]int
    numberMutex   sync.Mutex
    nextNumber    int
}

func NewConnectionManager() *ConnectionManager {
    return &ConnectionManager{
        onlineUsers:   make(map[string]*model.UserInfo),
        userNumberMap: make(map[string]int),
        nextNumber:    1,
    }
}

func (cm *ConnectionManager) AddUser(userID, ip string) int {
    cm.numberMutex.Lock()
    defer cm.numberMutex.Unlock()
    key := userID + "_" + ip
    if num, exists := cm.userNumberMap[key]; exists {
        return num
    }
    num := cm.nextNumber
    cm.userNumberMap[key] = num
    cm.nextNumber++
    return num
}

func (cm *ConnectionManager) UpdateOnline(userID string, info *model.UserInfo) {
    cm.onlineMutex.Lock()
    defer cm.onlineMutex.Unlock()
    info.LastSeen = time.Now().Unix()
    cm.onlineUsers[userID] = info
}

func (cm *ConnectionManager) RemoveUser(userID string) {
    cm.onlineMutex.Lock()
    defer cm.onlineMutex.Unlock()
    delete(cm.onlineUsers, userID)
}

func (cm *ConnectionManager) OnlineCount() int {
    cm.onlineMutex.RLock()
    defer cm.onlineMutex.RUnlock()
    return len(cm.onlineUsers)
}
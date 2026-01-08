package storage

import (
    "sync"
    "chat-server/model"
)

type MemoryStore struct {
    messages []model.ChatMessage
    mutex    sync.RWMutex
    maxSize  int
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{
        messages: make([]model.ChatMessage, 0, 120),
        maxSize:  120,
    }
}

func (ms *MemoryStore) Add(msg model.ChatMessage) {
    ms.mutex.Lock()
    defer ms.mutex.Unlock()
    ms.messages = append(ms.messages, msg)
    if len(ms.messages) > ms.maxSize {
        ms.messages = ms.messages[1:]
    }
}

func (ms *MemoryStore) GetAll() []model.ChatMessage {
    ms.mutex.RLock()
    defer ms.mutex.RUnlock()
    messages := make([]model.ChatMessage, len(ms.messages))
    copy(messages, ms.messages)
    return messages
}

func (ms *MemoryStore) Count() int {
    ms.mutex.RLock()
    defer ms.mutex.RUnlock()
    return len(ms.messages)
}
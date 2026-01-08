package middleware

import (
    "sync"
    "time"
)

type RateLimiter struct {
    ipMessages   map[string][]int64
    ipMutex      sync.RWMutex
    bannedIPs    map[string]int64
    banMutex     sync.RWMutex
}

func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        ipMessages: make(map[string][]int64),
        bannedIPs:  make(map[string]int64),
    }
}

func (rl *RateLimiter) CheckLimit(ip string) bool {
    rl.banMutex.RLock()
    if banUntil, exists := rl.bannedIPs[ip]; exists {
        if time.Now().Unix() < banUntil {
            rl.banMutex.RUnlock()
            return false
        }
    }
    rl.banMutex.RUnlock()
    rl.ipMutex.Lock()
    defer rl.ipMutex.Unlock()
    now := time.Now().Unix()
        rl.ipMessages[ip] = []int64{}
    }
    times := rl.ipMessages[ip]
    filtered := []int64{}
    for _, t := range times {
        if now-t < 60 {
            filtered = append(filtered, t)
        }
    }
    if len(filtered) >= 6 {
        rl.banMutex.Lock()
        rl.bannedIPs[ip] = now + 300
        rl.banMutex.Unlock()
        return false
    }
    filtered = append(filtered, now)
    rl.ipMessages[ip] = filtered
    return true
}
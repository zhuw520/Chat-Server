package utils

import (
    "strings"
    "net/http"
    "fmt"
)

func GetRealIP(headers http.Header, remoteAddr string) string {
    if xff := headers.Get("X-Forwarded-For"); xff != "" {
        ips := strings.Split(xff, ",")
        if len(ips) > 0 {
            return strings.TrimSpace(ips[0])
        }
    }
    if xRealIP := headers.Get("X-Real-IP"); xRealIP != "" {
        return strings.TrimSpace(xRealIP)
    }
    if strings.Contains(remoteAddr, ":") {
        parts := strings.Split(remoteAddr, ":")
        if len(parts) > 0 {
            return strings.Trim(parts[0], "[]")
        }
    }
    return remoteAddr
}

func IsPrivateIP(ip string) bool {
    if strings.Contains(ip, ":") {
        ip = strings.Split(ip, ":")[0]
        ip = strings.Trim(ip, "[]")
    }
    if strings.HasPrefix(ip, "192.168.") || 
       strings.HasPrefix(ip, "10.") || 
       ip == "127.0.0.1" {
        return true
    }
    if strings.HasPrefix(ip, "172.") {
        parts := strings.Split(ip, ".")
        if len(parts) >= 2 {
            second := 0
            fmt.Sscanf(parts[1], "%d", &second)
            return second >= 16 && second <= 31
        }
    }
    return false
}
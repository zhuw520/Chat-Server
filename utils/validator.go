package utils

import "strings"

func EscapeHTML(text string) string {
    text = strings.ReplaceAll(text, "&", "&amp;")
    text = strings.ReplaceAll(text, "<", "&lt;")
    text = strings.ReplaceAll(text, ">", "&gt;")
    text = strings.ReplaceAll(text, """, "&quot;")
    text = strings.ReplaceAll(text, "'", "&#39;")
    return text
}

func ValidateMessage(msg string) bool {
    msg = strings.TrimSpace(msg)
    if len(msg) == 0 {
        return false
    }
    if len(msg) > 78 {
        return false
    }
    return true
}
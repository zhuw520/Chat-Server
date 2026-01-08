package monitor

import (
    "sync"
    "time"
)

type StatsCollector struct {
    totalMessages   int64
    totalConnections int64
    startTime       time.Time
    mutex           sync.RWMutex
}

func NewStatsCollector() *StatsCollector {
    return &StatsCollector{
        startTime: time.Now(),
    }
}

func (sc *StatsCollector) RecordMessage() {
    sc.mutex.Lock()
    defer sc.mutex.Unlock()
    sc.totalMessages++
}

func (sc *StatsCollector) RecordConnection() {
    sc.mutex.Lock()
    defer sc.mutex.Unlock()
    sc.totalConnections++
}

func (sc *StatsCollector) GetStats() (int64, int64, time.Duration) {
    sc.mutex.RLock()
    defer sc.mutex.RUnlock()
    return sc.totalMessages, sc.totalConnections, time.Since(sc.startTime)
}
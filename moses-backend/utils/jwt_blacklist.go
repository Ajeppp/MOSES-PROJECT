package utils

import (
    "sync"
    "time"
)

type blacklistStore struct {
    mu   sync.RWMutex
    data map[string]int64 // token -> expiry unix
}

var bl *blacklistStore

func init() {
    bl = &blacklistStore{data: make(map[string]int64)}
    go bl.cleanupLoop()
}

func (b *blacklistStore) Add(token string, expUnix int64) {
    if token == "" {
        return
    }
    if expUnix <= time.Now().Unix() {
        return
    }
    b.mu.Lock()
    b.data[token] = expUnix
    b.mu.Unlock()
}

func (b *blacklistStore) IsBlacklisted(token string) bool {
    if token == "" {
        return false
    }
    b.mu.RLock()
    exp, ok := b.data[token]
    b.mu.RUnlock()
    if !ok {
        return false
    }
    if time.Now().Unix() > exp {
        b.mu.Lock()
        delete(b.data, token)
        b.mu.Unlock()
        return false
    }
    return true
}

func (b *blacklistStore) cleanupLoop() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()
    for range ticker.C {
        now := time.Now().Unix()
        b.mu.Lock()
        for t, e := range b.data {
            if e <= now {
                delete(b.data, t)
            }
        }
        b.mu.Unlock()
    }
}

// Exported helpers
func BlacklistAdd(token string, expUnix int64) {
    bl.Add(token, expUnix)
}

func BlacklistContains(token string) bool {
    return bl.IsBlacklisted(token)
}

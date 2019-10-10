package main

import (
	"fmt"
	"path/filepath"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Handler struct {
	mu         *sync.Mutex
	loggers    map[string]*lumberjack.Logger
	path       string
	maxSize    int
	maxAge     int
	maxBackups int
}

func (h *Handler) Handle(data map[string]interface{}) {
	h.logger(data["tag"].(string)).Write([]byte(data["content"].(string)))
}

func (h *Handler) logger(tag string) *lumberjack.Logger {
	h.mu.Lock()
	defer h.mu.Unlock()
	l, ok := h.loggers[tag]
	if ok {
		return l
	}

	l = &lumberjack.Logger{
		Filename:   filepath.Join(h.path, fmt.Sprintf("%s.log", tag)),
		MaxSize:    h.maxSize,
		MaxAge:     h.maxAge,
		MaxBackups: h.maxBackups,
		LocalTime:  true,
		Compress:   true,
	}
	h.loggers[tag] = l
	return l
}

func NewHandler(path string, maxSize, maxAge, maxBackups int) *Handler {
	return &Handler{
		mu:         new(sync.Mutex),
		loggers:    make(map[string]*lumberjack.Logger, 0),
		path:       path,
		maxSize:    maxSize,
		maxAge:     maxAge,
		maxBackups: maxBackups,
	}
}

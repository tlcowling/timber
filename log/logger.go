package log

import (
	"github.com/akerl/timber/log/levels"
)

// Logger carries template event fields and allows for logging events
type Logger struct {
	Fields map[string]string
}

// NewLogger returns a logger with a "name" field populated
func NewLogger(name string) *Logger {
	return &Logger{
		Fields: map[string]string{
			"name": name,
		},
	}
}

// Info logs an event at the INFO level
func (l *Logger) Info(fields map[string]string) {
	l.log(levels.LevelInfo, fields)
}

// Debug logs an event at the DEBUG level
func (l *Logger) Debug(fields map[string]string) {
	l.log(levels.LevelDebug, fields)
}

// InfoMsg logs an event at the INFO level with a string message
func (l *Logger) InfoMsg(msg string) {
	l.Info(map[string]string{"msg": msg})
}

// DebugMsg logs an event at the DEBUG level with a string message
func (l *Logger) DebugMsg(msg string) {
	l.Debug(map[string]string{"msg": msg})
}

func (l *Logger) log(lvl levels.Level, fields map[string]string) {
	globalProcessor.log(lvl, l.Fields, fields)
}

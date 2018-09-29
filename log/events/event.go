package events

import (
	"time"

	"github.com/akerl/timber/log/levels"
)

// Event describes a single log entry
type Event struct {
	Time   time.Time
	Fields map[string]string
}

// NewEvent creates a new empty Event whose timestamp is Now()
func NewEvent() *Event {
	return &Event{
		Time:   time.Now(),
		Fields: map[string]string{},
	}
}

// AddFields adds a set of fields to the event
func (e *Event) AddFields(f map[string]string) {
	for k, v := range f {
		e.Fields[k] = v
	}
}

// AddLevel adds a log level as a string to the event
func (e *Event) AddLevel(l levels.Level) {
	e.Fields["level"] = l.String()
}

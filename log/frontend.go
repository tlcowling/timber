package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	levelEnvVar     = "LOG_LEVEL"
	defaultLogLevel = "QUIET"
)

type event struct {
	Time   time.Time
	Fields map[string]string
}

func newEvent() *event {
	return &event{
		Time:   time.Now(),
		Fields: map[string]string{},
	}
}

func (e *event) AddFields(f fields) {
	for k, v := range f {
		e.Fields[k] = v
	}
}

func (e *event) AddLevel(l level) {
	e.Fields["level"] = l.String()
}

type level int

const (
	levelTrace level = iota + 1
	levelDebug
	levelInfo
	levelQuiet
)

var levelStrings = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"QUIET",
}

func (l level) String() string {
	if l < levelTrace || l > levelQuiet {
		return "UNKNOWN"
	}
	return levelStrings[l-1]
}

func levelFromString(str string) level {
	var index int
	for i, v := range levelStrings {
		if v == str {
			index = i
			break
		}
	}
	return level(index + 1)
}

type fields map[string]string

// Logger carries template event fields and allows for logging events
type Logger struct {
	Fields fields
}

// GetLogger returns a logger with a "name" field populated
func GetLogger(name string) *Logger {
	return &Logger{
		Fields: map[string]string{
			"name": name,
		},
	}
}

// Info logs an event at the INFO level
func (l *Logger) Info(f fields) {
	l.log(levelInfo, f)
}

// Debug logs an event at the DEBUG level
func (l *Logger) Debug(f fields) {
	l.log(levelDebug, f)
}

// Trace logs an event at the TRACE level
func (l *Logger) Trace(f fields) {
	l.log(levelTrace, f)
}

// InfoMsg logs an event at the INFO level with a string message
func (l *Logger) InfoMsg(msg string) {
	l.Info(map[string]string{"msg": msg})
}

// DebugMsg logs an event at the DEBUG level with a string message
func (l *Logger) DebugMsg(msg string) {
	l.Debug(map[string]string{"msg": msg})
}

// TraceMsg logs an event at the TRACE level with a string message
func (l *Logger) TraceMsg(msg string) {
	l.Trace(map[string]string{"msg": msg})
}

func (l *Logger) log(lvl level, f fields) {
	globalProcessor.log(lvl, l.Fields, f)
}

type processor struct {
	Level   level
	Encoder encoder
	Writer  writer
	Catcher catcher
}

type encoder interface {
	Encode(*event) (*string, error)
}

type writer interface {
	Write(*string) error
}

type catcher interface {
	Catch(error)
}

type stringEncoder struct{}

func (s stringEncoder) Encode(e *event) (*string, error) {
	var b strings.Builder
	var res string
	_, err := b.WriteString(e.Time.Format(time.RFC3339))
	if err != nil {
		return &res, err
	}
	for k, v := range e.Fields {
		_, err := fmt.Fprintf(&b, " %s=%s", k, v)
		if err != nil {
			return &res, err
		}
	}
	res = b.String()
	return &res, nil
}

type stderrWriter struct{}

func (s stderrWriter) Write(str *string) error {
	_, err := os.Stderr.WriteString(*str)
	return err
}

type panicCatcher struct{}

func (p panicCatcher) Catch(err error) {
	panic(err)
}

var globalProcessor processor

func init() {
	envLevelStr := os.Getenv(levelEnvVar)
	if envLevelStr == "" {
		envLevelStr = defaultLogLevel
	}

	envLevel := levelFromString(envLevelStr)

	globalProcessor = processor{
		Level:   envLevel,
		Encoder: stringEncoder{},
		Writer:  stderrWriter{},
		Catcher: panicCatcher{},
	}
}

func (p *processor) log(lvl level, fl ...fields) {
	if lvl < p.Level {
		return
	}
	e := newEvent()
	for _, f := range fl {
		e.AddFields(f)
	}
	e.AddLevel(lvl)
	str, err := p.Encoder.Encode(e)
	if err != nil {
		p.Catcher.Catch(err)
	}
	err = p.Writer.Write(str)
	if err != nil {
		p.Catcher.Catch(err)
	}
}

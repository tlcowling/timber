package log

import (
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
	return &Event{
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

type logger struct {
	Fields fields
}

// GetLogger returns a logger with a "name" field populated
func GetLogger(name string) *logger {
	return &logger{
		Fields: map[string]string{
			"name": name,
		},
	}
}

func (l *logger) Info(f fields) {
	l.log(levelInfo, f)
}

func (l *logger) Debug(f fields) {
	l.log(levelDebug, f)
}

func (l *logger) Trace(f fields) {
	l.log(levelTrace, f)
}

func (l *logger) log(lvl level, f fields) {
	globalProcessor.log(lvl, l.Fields, f)
}

type processor struct {
	Level   level
	Encoder encoder
	Writer  writer
	OnFail  errorHandler
}

type encoder interface {
	Encode(*event) (*string, error)
}

type writer interface {
	Write(*string) error
}

type errorHandler interface {
	Catch(error)
}

type stringEncoder struct{}

func (s stringEncoder) Encode(e *event) (*string, error) {
	var b strings.Builder
	_, err := b.WriteString(t.Format(time.RFC3339))
	if err != nil {
		return "", err
	}
	for k, v := range e.Fields {
		_, err := fmt.Fprintf(&b, " %s=%s", k, v)
		if err != nil {
			return "", err
		}
	}
	return &b.String(), nil
}

type stderrWriter struct{}

func (s stderrWriter) Write(str *string) error {
	_, err := os.Stderr.WriteString(string)
	return err
}

type panicHandler struct{}

func (p panicHandler) Catch(err error) {
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
		OnFail:  panicHandler{},
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
	str, err := p.Encoder.Encode(&e)
	if err != nil {
		p.OnFail(err)
	}
	err = p.Writer.Write(str)
	if err != nil {
		p.OnFail(err)
	}
}

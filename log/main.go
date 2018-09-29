package log

type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
)

type Encoder interface {
	Encode(*Event) ([]byte, error)
}

type Writer interface {
	Write(*Event)
}

type config struct {
	Level   Level
	Encoder Encoder
	Writer  Writer
}

type Event struct {
	Time   time.Time
	Fields map[string]interface{}
}

type Field struct {
	Name  string
	Value interface{}
}

func NewEvent(fields ...Field) *Event {
	return &Event{
		Time:   time.Now(),
		Fields: fields,
	}
}

package levels

// Level describes a log verbosity level
type Level int

const (
	// LevelQuiet disables all logging
	LevelQuiet Level = iota
	// LevelInfo shows information about process flow
	LevelInfo
	// LevelDebug shows low-level information about process actions
	LevelDebug
)

var levelStrings = []string{
	"QUIET",
	"INFO",
	"DEBUG",
}

// String returns the text representation of the log level
func (l Level) String() string {
	if l < LevelQuiet || l > LevelDebug {
		return "UNKNOWN"
	}
	return levelStrings[l]
}

// FromString parses a log level from a name
func FromString(str string) Level {
	var index int
	for i, v := range levelStrings {
		if v == str {
			index = i
			break
		}
	}
	return Level(index)
}

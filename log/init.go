package log

import (
	"os"

	"github.com/akerl/timber/log/levels"
)

const (
	levelEnvVar     = "LOG_LEVEL"
	defaultLogLevel = levels.LevelQuiet
)

var globalProcessor processor

func init() {
	l := defaultLogLevel
	if envLevel, found := os.LookupEnv(levelEnvVar); found {
		l = levels.FromString(envLevel)
	}

	globalProcessor = processor{
		Level:   l,
		Encoder: stringEncoder{},
		Writer:  stderrWriter{},
		Catcher: panicCatcher{},
	}
}

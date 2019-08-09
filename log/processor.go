package log

import (
	"github.com/akerl/timber/v2/log/events"
	"github.com/akerl/timber/v2/log/levels"
)

type processor struct {
	Level   levels.Level
	Encoder encoder
	Writer  writer
	Catcher catcher
}

type encoder interface {
	Encode(events.Event) (string, error)
}

type writer interface {
	Write(string) error
}

type catcher interface {
	Catch(error)
}

func (p processor) log(lvl levels.Level, fl ...map[string]string) {
	if lvl > p.Level {
		return
	}
	e := events.NewEvent()
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

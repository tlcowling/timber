package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/akerl/timber/log/events"
)

type stringEncoder struct{}

func (s stringEncoder) Encode(e *events.Event) (*string, error) {
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

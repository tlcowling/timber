package log

import (
	"os"
)

type stderrWriter struct{}

func (s stderrWriter) Write(str *string) error {
	_, err := os.Stderr.WriteString(*str)
	return err
}

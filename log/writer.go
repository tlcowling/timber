package log

import (
	"fmt"
	"os"
)

type stderrWriter struct{}

func (s stderrWriter) Write(str *string) error {
	_, err := fmt.Fprintln(os.Stderr, *str)
	return err
}

package logger

import (
	"fmt"
	"os"
)

type Closer struct {
	Logger Logger
	file   *os.File
}

func (l Closer) Close() error {
	_ = l.Logger.Sync()

	if l.file != nil {
		err := l.file.Close()
		if err != nil {
			return fmt.Errorf("close file: %w", err)
		}
	}

	return nil
}

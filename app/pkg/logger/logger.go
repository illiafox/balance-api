package logger

import (
	"fmt"
	"os"

	zapcore "balance-service/app/pkg/logger/zap"
)

const separator = "\n\n"

func New(path string) (Closer, error) {
	var logger Closer

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return logger, fmt.Errorf("create/open log file (%s): %w", path, err)
	}

	info, err := file.Stat()
	if err != nil {
		return logger, fmt.Errorf("get file stats: %w", err)
	}

	if info.IsDir() {
		return logger, fmt.Errorf("%s is directory", info.Name())
	}

	if info.Size() > 0 {
		_, err = file.WriteString(separator)
		if err != nil {
			return logger, fmt.Errorf("write separator to file: %w", err)
		}
	}

	logger.Logger = zapcore.NewLogger(os.Stdout, file)
	logger.file = file

	return logger, nil
}

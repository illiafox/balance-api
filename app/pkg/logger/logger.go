package logger

import (
	"fmt"
	"os"

	zapcore "balance-service/app/pkg/logger/zap"
	"go.uber.org/zap"
)

const separator = "\n\n"

func New(path string) (*zap.Logger, func() error, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("create/open log file (%s): %w", path, err)
	}

	info, err := file.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("get file stats: %w", err)
	}

	if info.IsDir() {
		return nil, nil, fmt.Errorf("%s is directory", info.Name())
	}

	if info.Size() > 0 {
		_, err = file.WriteString(separator)
		if err != nil {
			return nil, nil, fmt.Errorf("write separator to file: %w", err)
		}
	}

	return zapcore.NewLogger(os.Stdout, file), file.Close, nil
}

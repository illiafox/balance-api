package logger

import (
	"fmt"
	"os"

	zapcore "balance-service/app/pkg/logger/zap"
	"go.uber.org/zap"
)

const separator = "\n\n"

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Named(s string) *zap.Logger
}

type Zap struct {
	Logger *zap.Logger
	file   *os.File
}

func (l Zap) Close() error {
	_ = l.Logger.Sync()

	if l.file != nil {
		err := l.file.Close()
		if err != nil {
			return fmt.Errorf("close file: %w", err)
		}
	}

	return nil
}

func New(path string) (Zap, error) {
	var logger Zap

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

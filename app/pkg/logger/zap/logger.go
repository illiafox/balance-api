package zap

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func (l Logger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l Logger) Named(s string) Logger {
	return Logger{l.logger.Named(s)}
}

func (l Logger) With(fields ...Field) Logger {
	return Logger{l.logger.With(fields...)}
}

func (l Logger) Sync() error {
	return l.logger.Sync()
}

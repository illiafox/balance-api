package composites

import "go.uber.org/zap"

type LoggerComposite struct {
	logger *zap.Logger
}

func NewLoggerComposite(logger *zap.Logger) LoggerComposite {
	return LoggerComposite{logger: logger}
}

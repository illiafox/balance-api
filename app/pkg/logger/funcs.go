package logger

import "balance-service/app/pkg/logger/zap"

func Bool(key string, value bool) Field {
	return zap.Bool(key, value)
}

func Int64(key string, value int64) Field {
	return zap.Int64(key, value)
}

func Any(key string, value any) Field {
	return zap.Any(key, value)
}

func Error(value error) Field {
	return zap.Error(value)
}

func String(key string, value string) Field {
	return zap.String(key, value)
}

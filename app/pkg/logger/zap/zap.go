package zap

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(console io.Writer, files ...io.Writer) *zap.Logger {
	pe := zap.NewProductionEncoderConfig()
	// file
	pe.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)
	fileEncoder := zapcore.NewJSONEncoder(pe)
	// console
	pe.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05") // "02/01/2006 15:04:05 |"
	pe.EncodeLevel = levelEncoder
	pe.ConsoleSeparator = " "
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	// //
	cores := make([]zapcore.Core, len(files)+1)

	// console
	cores[0] = zapcore.NewCore(consoleEncoder,
		zapcore.AddSync(console),
		zap.DebugLevel,
	)

	// // add syncers
	for i := range files {
		cores[i+1] = zapcore.NewCore(fileEncoder,
			zapcore.AddSync(files[i]),
			zap.DebugLevel,
		)
	}

	//
	return zap.New(
		zapcore.NewTee(cores...),
	)
}

func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("|")
	enc.AppendString(l.CapitalString())
	enc.AppendString("|")
}

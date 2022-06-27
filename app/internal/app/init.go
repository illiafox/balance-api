package app

import (
	"flag"
	"log"
	"runtime"

	"balance-service/app/pkg/closer"
	logs "balance-service/app/pkg/logger"
)

type flags struct {
	https, noswag bool
}

func New() *App {

	var (
		logPath = flag.String("log", "log.txt", "log file path")
		//
		https  = flag.Bool("https", false, "run server in https mode")
		noswag = flag.Bool("noswag", false, "disable swagger")
	)

	flag.Parse()

	// // logger
	logger, err := logs.New(*logPath)
	if err != nil {
		log.Fatalf("init logger: %v", err)
	}

	defer runtime.GC() // force garbage collector to clear unused flag pointers

	app := App{
		flags: flags{
			https:  *https,
			noswag: *noswag,
		},
		//
		logger: logger.Logger,
		//
		closers: closer.New(logger.Logger),
	}

	// add logger close
	app.closers.Add(logger)

	return &app
}

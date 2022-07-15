package app

import (
	"flag"
	"log"
	"runtime"

	logs "balance-service/app/pkg/logger"
)

type flags struct {
	https, swagger, pprof, prom bool
}

func New() *App {

	var (
		logPath = flag.String("log", "log.txt", "log file path")
		//
		https  = flag.Bool("https", false, "run server in https mode")
		noswag = flag.Bool("swagger", false, "enable swagger")
		pprof  = flag.Bool("pprof", false, "enable pprof")
		prom   = flag.Bool("prom", false, "enable prometheus")
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
			https:   *https,
			swagger: *noswag,
			pprof:   *pprof,
			prom:    *prom,
		},
		//
		logger: logger.Logger,
	}

	// add logger close
	app.closers.Add(logger)

	return &app
}

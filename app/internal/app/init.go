package app

import (
	"flag"
	"log"
	"runtime"

	logs "balance-service/app/pkg/logger"
)

type flags struct {
	config string
	https  bool
}

func Init() App {

	var (
		logPath = flag.String("log", "log.txt", "log file path (default 'log.txt')")
		config  = flag.String("config", "config.toml", "config path (default 'config.toml')")
		//cache   = flag.Bool("cache", false, "use built-it storage")
		//debug   = flag.Bool("debug", false, "enable debug mode")
		https = flag.Bool("https", false, "run server in https mode")
	)

	flag.Parse()

	// // logger
	logger, closer, err := logs.New(*logPath)
	if err != nil {
		log.Fatalf("init logger: %v", err)
	}

	defer runtime.GC() // force garbage collector to clear unused flag pointers

	app := App{
		flags: flags{
			//cache:  *cache,
			config: *config,
			//debug:  *debug,
			https: *https,
		},

		logger: logger,
	}
	app.closers.logger = closer

	return app
}

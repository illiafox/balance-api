package main

import (
	apps "balance-service/app/internal/app"
)

// @title Balance API
// @version 1.0
// @description Balance API.
// @termsOfService https://swagger.io/terms/

// @contact.name Developer
// @contact.url https://github.com/illiafox
// @contact.email illiadimura@gmail.com

// @license.name Boost Software License 1.0
// @license.url https://opensource.org/licenses/BSL-1.0

// @host 0.0.0.0:8080
// @BasePath /api
// @schemes http https

func main() {
	app := apps.New()
	app.Run()
}

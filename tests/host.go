package tests

import (
	"net/url"
	"os"
	"strings"
)

var Host = "http://localhost:8080"

func init() {
	if host, ok := os.LookupEnv("HOST"); ok {
		// parse url
		u, err := url.Parse(host)
		if err != nil {
			panic(err)
		}

		// remove query
		u.RawQuery = ""

		// trim slash
		Host = strings.TrimRight(u.String(), "/")
	}
}

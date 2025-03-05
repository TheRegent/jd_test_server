package main

import (
	"net/http"
	"os"

	"jd_test_server/src/handler"
)

const (
	DefaultPort = "8080"
	PortEnvVar  = "JD_TEST_PORT"
)

func main() {
	handler.SetupRoutes()

	port, ok := os.LookupEnv(PortEnvVar)
	if !ok {
		port = DefaultPort
	}

	http.ListenAndServe(":"+port, nil)
}

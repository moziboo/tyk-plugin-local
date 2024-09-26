package middleware

import (
	"net/http"
	"plugin-dev/util/logger"
)

func ResponseMiddleware(rw http.ResponseWriter, res *http.Response, req *http.Request) {
	logger.Info("--- Required Client header! ----")
}

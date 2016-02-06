package gorillalogger

import (
	"github.com/gorilla/handlers"
	"github.com/srinathh/middleware"
	"io"
	"net/http"
)

// GorillaLogger wraps the handlers.LoggingHandler from gorilla framework to gennerate
// standard Apache Logs of the request and response
func GorillaLogger(w io.Writer) middleware.MiddleWare {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(w, next)
	}
}

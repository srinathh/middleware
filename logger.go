package middleware

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
)

// GorillaLogger wraps handlers.LoggingHandler from the gorilla toolkit
type GorillaLogger struct {
	io.Writer
}

// Log is the middleware function wrapping Gorilla's logging handler
func (g GorillaLogger) Log(next http.Handler) http.Handler {
	return handlers.LoggingHandler(g, next)
}

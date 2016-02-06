package logging

import (
	"github.com/srinathh/middleware"
	"log"
	"net/http"
	"time"
)

// LogLogger logs requests to the standard library Go logger. Unlike implementations of
// the full Apache logging format, this will not read the output to determine return
// code or bytes written to the response writer and should have less overhead
func LogLogger() middleware.MiddleWare {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s %s", r.RemoteAddr, time.Now().Format(time.RFC3339), r.Method, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}

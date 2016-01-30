// This file defines some small and simple middleware

package middleware

import (
	"fmt"
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"strings"
)

// MiddleWare is a function that takes a http.Handler and wraps your middleware around it.
// Call next.ServeHTTP in your middleware if you want to continue the middleware chain
// or return to stop propagation.
type MiddleWare func(next http.Handler) http.Handler

// GorillaLogger wraps the handlers.LoggingHandler from gorilla framework to gennerate
// standard Apache Logs of the request and response
func GorillaLogger(w io.Writer) MiddleWare {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(w, next)
	}
}

// CannedResponse middleware detects URL.Path starting wit a given pattern and
// returns a canned response. Non matching responses are passed on to the next
// handler in the middleware chain
func CannedResponse(pattern, response string) MiddleWare {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Index(r.URL.Path, pattern) == 0 {
				fmt.Fprint(w, response)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

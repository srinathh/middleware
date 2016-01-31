// This file defines some small and simple middleware

package middleware

import (
	"fmt"
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"regexp"
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

// Intercept is a convenience Filter middleware that blocks a specific pattern
// of URL.String() and resonds with provided response string
func Intercept(pattern, response string) MiddleWare {
	return Filter(
		true,
		[]string{pattern},
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, response)
		}),
	)
}

func makeRegexps(patterns []string) []*regexp.Regexp {
	ret := make([]*regexp.Regexp, len(patterns))
	for _, txt := range patterns {
		ret = append(ret, regexp.MustCompile(txt))
	}
	return ret
}

// Filter whitelists or blacklists URL.String() matched the provide regex patterns
// Whitelisting passes through only matched requests and Blacklisting passes through
// non-matched requests. blockHandler is called for the blocked requests.
func Filter(blacklist bool, patterns []string, blockHandler http.Handler) MiddleWare {
	regxs := makeRegexps(patterns)
	return func(next http.Handler) http.Handler {
		var inLoop, outLoop http.Handler
		if blacklist {
			inLoop = blockHandler
			outLoop = next
		} else {
			inLoop = next
			outLoop = blockHandler
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			urlstr := r.URL.String()
			for _, regx := range regxs {
				if regx.MatchString(urlstr) {
					inLoop.ServeHTTP(w, r)
					return
				}
			}
			outLoop.ServeHTTP(w, r)
		})
	}
}

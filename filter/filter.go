package filter

import (
	"fmt"
	"github.com/srinathh/middleware"
	"net/http"
	"regexp"
)

// Intercept is a convenience Filter middleware that blocks a specific pattern
// of URL.String() and resonds with provided response string
func Intercept(pattern, response string) middleware.MiddleWare {
	return Filter(
		true,
		[]string{pattern},
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, response)
		}),
	)
}

// Filter whitelists or blacklists URL.String() matched the provide regex patterns
// Whitelisting passes through only matched requests and Blacklisting passes through
// non-matched requests. blockHandler is called for the blocked requests.
func Filter(blacklist bool, patterns []string, blockHandler http.Handler) middleware.MiddleWare {
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

			for _, pattern := range patterns {
				matched, err := regexp.MatchString(pattern, urlstr)
				if err != nil {
					panic(err)
				}
				if matched {
					inLoop.ServeHTTP(w, r)
					return
				}
			}
			outLoop.ServeHTTP(w, r)
		})
	}
}

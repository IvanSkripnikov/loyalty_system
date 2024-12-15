package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"loyalty_system/controllers"
	"loyalty_system/logger"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
}

func initHTTPServer() {
	http.HandleFunc("/", Serve)

	err := http.ListenAndServe(":8080", nil) //nolint:gosec
	if err != nil {
		errMessage := fmt.Sprintf("Can't init HTTP server: %v", err)

		logger.Error(errMessage)
	}
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string

	found := false

	for _, routeUnit := range routes {
		matches := routeUnit.regex.FindStringSubmatch(strings.TrimSpace(r.URL.Path))

		if len(matches) > 0 {
			if r.Method != routeUnit.method {
				allow = append(allow, routeUnit.method)

				continue
			}

			found = true

			routeUnit.handler(w, r)
		}
	}

	if !found && len(allow) == 0 {
		w.WriteHeader(http.StatusNotFound)

		http.NotFound(w, r)

		return
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))

		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)

		return
	}
}

func GetHTTPHandler() *http.ServeMux {
	httpHandler := http.NewServeMux()

	for _, routeUnit := range routes {
		httpHandler.HandleFunc(handleRegexp(routeUnit.regex), routeUnit.handler)
	}

	return httpHandler
}

func handleRegexp(regExp *regexp.Regexp) string {
	expr := regExp.String()[1 : len(regExp.String())-1]

	var result string

	if strings.Count(expr, "/") > 1 {
		parts := strings.Split(expr, "/")

		parts = parts[:len(parts)-1]

		result = strings.Join(parts, "/") + "/"
	} else {
		result = expr
	}

	return result
}

func main() {
	initHTTPServer()
}

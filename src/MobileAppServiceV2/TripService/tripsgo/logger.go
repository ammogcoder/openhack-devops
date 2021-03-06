package tripsgo

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Logger - basic console logger that writes request info to stdout
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func LogToConsole(message string) {
	var debug, present = os.LookupEnv("DEBUG_LOGGING")

	if present && debug == "true" {
		log.Printf(message)
	}
}

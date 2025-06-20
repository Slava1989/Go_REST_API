package http

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w, r)
	})
}

func LoggindMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.WithFields(
		// 	log.Fields(
		// 		"method": r.Method,
		// 		"path": r.URL.Path,
		// 	)).Info("handled request")
		// next.ServeHTTP(w, r)
		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("handled request")
		next.ServeHTTP(w, r)
	})
}

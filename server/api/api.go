package api

import (
	"net/http"
	"github.com/go-chi/chi/v5"

	"mc.honki.ng/spwnr/logging"

	"mc.honki.ng/spwnr/api/routes/ping"
)

func CreateApiRouter(logger logging.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(loggerMiddleware(logger))

	r.Get("/ping", ping.Ping)

	return r
}

func loggerMiddleware(logger logging.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Log(r.Method + " " + r.URL.Path + " from " + r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	}
}

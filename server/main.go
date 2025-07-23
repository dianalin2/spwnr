package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"mc.honki.ng/spwnr/logging"
	"mc.honki.ng/spwnr/api"
)

func createRoutes(logger logging.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/api/v1", api.CreateApiRouter(logger))

	return r
}


func main() {
	logger, err := logging.CreateFileLogger("server.log")
	addr := ":8080"
	if err != nil {
		panic("Failed to create logger: " + err.Error())
	}

	server := &http.Server{
		Addr:    addr,
		Handler: createRoutes(logger),
	}

	// Start the server
	logger.Log("Starting server on port " + addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log("Server failed to start: " + err.Error())
		return
	}

	defer func() {
		if err := server.Close(); err != nil {
			logger.Log("Server failed to close: " + err.Error())
		} else {
			logger.Log("Server closed successfully")
		}
	} ()
}

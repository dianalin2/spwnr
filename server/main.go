package main

import (
	"net/http"
	"mc.honki.ng/spwnr/logging"
	"mc.honki.ng/spwnr/api"
	"github.com/go-chi/chi/v5"
)

func createRoutes(logger logging.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/api/v1", api.CreateApiRouter(logger))

	return r
}


func main() {
	logger, err := logging.CreateFileLogger("server.log")
	if err != nil {
		panic("Failed to create logger: " + err.Error())
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: createRoutes(logger),
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log("Server failed to start: " + err.Error())
		return
	}

	logger.Log("Server started successfully")

	defer func() {
		if err := server.Close(); err != nil {
			logger.Log("Server failed to close: " + err.Error())
		} else {
			logger.Log("Server closed successfully")
		}
	} ()
}

package servers

import (
	"net/http"
	"encoding/json"

	"github.com/go-chi/chi/v5"
)

func Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/info", getInfo)

	return router
}

type ServerResponse struct {
	Message string
	Error string `json:"error,omitempty"`
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	response := ServerResponse {Message: "Development Version: 0.0.0"}

	w.WriteHeader(http.StatusOK)
	error := json.NewEncoder(w).Encode(response)
	if error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

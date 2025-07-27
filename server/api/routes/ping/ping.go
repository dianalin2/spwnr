package ping

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Init() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", Ping)

	return router
}

type PingMessage struct {
	Status string `json:"status"`
	Message string `json:"message,omitempty"`
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := PingMessage{
		Status: "ok",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

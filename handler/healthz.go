package handler

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = &model.HealthzResponse{}
	healthz := model.HealthzResponse{
		Message: "OK",
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(healthz); err != nil {
		log.Fatal(err)
	}
}

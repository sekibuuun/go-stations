package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req model.CreateTODORequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		todo, err := h.svc.CreateTODO(context.Background(), req.Subject, req.Description)

		if req.Description == "" {
			w.WriteHeader(http.StatusOK)
		}

		if err != nil {
			log.Fatal(err)
		}
		res := model.CreateTODOResponse{
			TODO: *todo,
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(res); err != nil {
			log.Fatal(err)
		}
	}

	if r.Method == http.MethodPut {
		var req model.UpdateTODORequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatal(err)
			return
		}

		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		todo, err := h.svc.UpdateTODO(context.Background(), req.ID, req.Subject, req.Description)

		if todo == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err != nil {
			log.Fatal(err)
		}
		res := model.UpdateTODOResponse{
			TODO: *todo,
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(res); err != nil {
			log.Fatal(err)
		}
	}

	if r.Method == http.MethodGet {
		var req model.ReadTODORequest
		queryParams := r.URL.Query()

		if PrevID := queryParams.Get("prev_id"); PrevID != "" {
			req.PrevID, _ = strconv.ParseInt(PrevID, 10, 64)
		}
		if Size := queryParams.Get("size"); Size != "" {
			req.Size, _ = strconv.ParseInt(Size, 10, 64)
		}
		if req.Size == 0 {
			req.Size = 5
		}

		todos, err := h.svc.ReadTODO(context.Background(), req.PrevID, req.Size)

		if err != nil {
			log.Fatal(err)
		}
		res := model.ReadTODOResponse{
			TODOs: todos,
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(res); err != nil {
			log.Fatal(err)
		}
	}

	if r.Method == http.MethodDelete {
		var req model.DeleteTODORequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error decoding request:", err)
			return
		}

		if len(req.IDs) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Empty IDs list")
			return
		}

		err := h.svc.DeleteTODO(context.Background(), req.IDs)
		if err != nil {
			if _, ok := err.(*model.ErrNotFound); ok {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			log.Println("Error deleting TODO:", err)
			return
		}

		res := model.DeleteTODOResponse{}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error encoding response:", err)
			return
		}
		return
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

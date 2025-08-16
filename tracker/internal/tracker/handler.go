package tracker

import (
	respond "gogym/tracker/internal/shared"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TrackerHandler struct {
}

func NewTrackerHandler() *TrackerHandler {
	return &TrackerHandler{}
}

func (h *TrackerHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.List)
	return r
}

func (h *TrackerHandler) List(w http.ResponseWriter, r *http.Request) {
	trackers := []TrackerDto{}
	trackers = append(trackers, TrackerDto{ID: "1", Name: "Running"})
	respond.JSON(w, http.StatusOK, trackers)
}

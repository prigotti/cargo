package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prigotti/cargo/clientapi/application"
	"github.com/prigotti/cargo/common/pb"
)

type PortHTTPHandler struct {
	s application.PortService
}

// NewPortHTTPHandler is the Port HTTP API handler factory
// Receiving the router is a personal preference, as I enjoy
// configuring paths closer to the handler function definitions.
func NewPortHTTPHandler(r *mux.Router, s application.PortService) *PortHTTPHandler {
	h := &PortHTTPHandler{s: s}

	h.registerRoutes(r)

	return h
}

func (h *PortHTTPHandler) registerRoutes(r *mux.Router) {
	r.Methods("GET").Path("/ports").HandlerFunc(h.List)
}

// List will get a list of items from the domain service.
// Since the list might get to large, we enforce pagination
// with query parameters (when omitted, default values apply).
func (h *PortHTTPHandler) List(rw http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")
	perPageStr := req.URL.Query().Get("perPage")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = application.DefaultPage
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		perPage = application.DefaultPerPage
	}

	result, err := h.s.List(req.Context(), &pb.ListQuery{Page: uint32(page), PerPage: uint32(perPage)})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(encoded)
}

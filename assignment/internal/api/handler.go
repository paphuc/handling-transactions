package api

import (
	"encoding/json"
	"log"
	"net/http"

	"handling-transactions/assignment/pkg/http/response"
)

func NewHandler(srv ServiceI) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) AddAssignment(w http.ResponseWriter, r *http.Request) {
	log.Println("===== Add assignment =====")
	var o InsertAssignmentRequest
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	_, err := h.srv.InsertAssignment(r.Context(), o)
	if err != nil {
		log.Println("===== Addding assignment failed =====: ", err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, response.Base{
		ID: o.Body.ID,
	})
	log.Println("===== Added assignment =====: ", o.Body.ID)
}

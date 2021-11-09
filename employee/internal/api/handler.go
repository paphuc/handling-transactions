package api

import (
	"encoding/json"
	"log"
	"net/http"

	"handling-transactions/employee/pkg/http/response"
)

func NewHandler(srv ServiceI) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	log.Println("==== Insert Employee ====")
	var e EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	id, err := h.srv.InsertEmployee(r.Context(), e)
	if err != nil {
		log.Println("==== Added Employee failed ====", err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{
		"id": *id,
	})
	log.Println("==== Added Employee ====", *id)
}

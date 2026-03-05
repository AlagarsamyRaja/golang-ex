package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gorm/models"

	"gorm/service"
	"net/http"
)

type Handler struct {
	handl service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		handl: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.UserRoles

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.handl.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Header().Set("Content-Type", "application/json")

		err := models.Error{
			Message:    err.Error(),
			Errors:     err,
			Statuscode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user []models.UserRoles

	result, err := h.handl.Read(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Header().Set("Content-Type", "application/json")

		err := models.Error{
			Message:    err.Error(),
			Errors:     err,
			Statuscode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusMethodNotAllowed)
		return
	}

	var user models.UserRoles

	result, err := h.handl.ReadById(user, userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")

		err := models.Error{
			Message:    err.Error(),
			Errors:     err,
			Statuscode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updatedUserData models.UserRoles

	err = json.NewDecoder(r.Body).Decode(&updatedUserData)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = h.handl.UpdateById(updatedUserData, userId)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")

		err := models.Error{
			Message:    err.Error(),
			Errors:     err,
			Statuscode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Updated successfully")

}

func (h *Handler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var user models.UserRoles

	_, err = h.handl.DeleteById(user, userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")

		err := models.Error{
			Message:    err.Error(),
			Errors:     err,
			Statuscode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Deleted successfully")

}

func (h *Handler) PatchByid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updatedUserData models.UserRoles
	err = json.NewDecoder(r.Body).Decode(&updatedUserData)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = h.handl.PatchById(updatedUserData, userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")

		err := models.Error{
			Message:    err.Error(),
			Errors:     err,
			Statuscode: http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Updated successfully")
}


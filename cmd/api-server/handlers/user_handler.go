package handlers

import (
	"encoding/json"
	"net/http"

	"rush-free-server/internal/repository/postgres"

	"go.uber.org/zap"
)

type UserHandler struct {
	Repo   *postgres.UserRepository
	Logger *zap.SugaredLogger
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(repo *postgres.UserRepository, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{Repo: repo, Logger: logger}
}

// GetUsersHandler handles the GET /users endpoint.
func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetUsers()
	if err != nil {
		h.Logger.Errorf("Error fetching users: %v", err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		h.Logger.Errorf("Error encoding users response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetUserByNameHandler handles the GET /users/:name endpoint.
func (h *UserHandler) GetUsersByNameHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetUsersByName(r.URL.Path[7:])
	if err != nil {
		h.Logger.Errorf("Error fetching users: %v", err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		h.Logger.Errorf("Error encoding user response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

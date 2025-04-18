package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"goGate/internal/auth/service"
)

type Handler struct {
	authService service.AuthService
}

func NewHandler(authService service.AuthService) *Handler {
	return &Handler{authService: authService}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "неверный запрос", http.StatusBadRequest)
		return
	}
	token, err := h.authService.Register(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "неверный запрос", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "токен не предоставлен", http.StatusUnauthorized)
		return
	}
	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		http.Error(w, "неправильный формат токена", http.StatusUnauthorized)
		return
	}
	tokenStr := parts[1]

	_, err := h.authService.ValidateToken(tokenStr)
	if err != nil {
		http.Error(w, "недействительный токен", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Добро пожаловать!"))
}

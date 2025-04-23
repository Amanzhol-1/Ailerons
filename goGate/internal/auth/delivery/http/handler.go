package http

import (
	"net/http"

	"goGate/internal/auth/domain"
	"goGate/internal/auth/middleware"
	"goGate/internal/auth/service"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	authSvc service.AuthService
	profSvc service.ProfileService
}

func NewHandler(a service.AuthService, p service.ProfileService) *Handler {
	return &Handler{authSvc: a, profSvc: p}
}

func (h *Handler) Register(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	token, err := h.authSvc.Register(req.Username, req.Password, req.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) Login(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	token, err := h.authSvc.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) GetProfile(c echo.Context) error {
	t := c.Get("user").(*jwt.Token)
	claims := t.Claims.(*middleware.Claims)
	profile, err := h.profSvc.GetMyProfile(claims.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, profile)
}

func (h *Handler) UpdateCustomer(c echo.Context) error {
	t := c.Get("user").(*jwt.Token)
	claims := t.Claims.(*middleware.Claims)
	if claims.Role != "customer" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}
	var p domain.CustomerProfile
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid"})
	}
	if err := h.profSvc.UpdateCustomerProfile(claims.Username, &p); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handler) UpdateDriver(c echo.Context) error {
	t := c.Get("user").(*jwt.Token)
	claims := t.Claims.(*middleware.Claims)
	if claims.Role != "driver" {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "forbidden"})
	}
	var p domain.DriverProfile
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid"})
	}
	if err := h.profSvc.UpdateDriverProfile(claims.Username, &p); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, p)
}

package handlers

import (
	"crabi_test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler represents the API handler that uses the Service dependency.
type Handler struct {
	service service.Service
}

// NewHandler creates a new instance of the Handler.
func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

// GetUser is a handler function that retrieves data using the Service dependency.
func (h *Handler) GetUser(c *gin.Context) {
	data := h.service.FetchUser()
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) CreateUser(c *gin.Context) {
	data := h.service.CreateUser()
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) AuthUser(c *gin.Context) {
	data := h.service.AuthUser()
	c.JSON(http.StatusOK, gin.H{"data": data})
}

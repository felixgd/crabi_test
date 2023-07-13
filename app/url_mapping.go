package app

import (
	"crabi_test/handlers"

	"github.com/gin-gonic/gin"
)

// Router represents the API router that registers routes.
type Router struct {
	engine  *gin.Engine
	handler *handlers.Handler
}

// NewRouter creates a new instance of the Router.
func NewRouter(e *gin.Engine, h *handlers.Handler) *Router {
	return &Router{
		engine:  e,
		handler: h,
	}
}

// RegisterRoutes registers the API routes.
func (r *Router) RegisterRoutes() {
	r.engine.GET("/user/:ID", r.handler.GetUser)
	r.engine.POST("/user", r.handler.CreateUser)
	r.engine.POST("/auth", r.handler.AuthUser)
}

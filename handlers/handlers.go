package handlers

import (
	"crabi_test/domain"
	"crabi_test/service"
	"crabi_test/utils/constants"
	"crabi_test/utils/jwt"
	"log"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// GetUser is a handler function that retrieves data using the Service dependency.
func (h *Handler) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	token := ctx.GetHeader("auth")

	_, err := mail.ParseAddress(email)
	if err != nil {
		ctx.JSON(400, "Invalid Data")
		return
	}

	tokenEmail, err := jwt.VerifyToken(token, []byte(constants.SECRET_KEY))
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	if tokenEmail != email {
		ctx.JSON(http.StatusForbidden, "Not Authorized")
		return
	}

	data, err := h.service.FetchUser(ctx, email)
	if err != nil {
		ctx.JSON(400, nil)
		return
	}

	data.Password = ""

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := domain.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Print("error on bindJSON")
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}

	err := validate.Struct(user)
	if err != nil {
		ctx.JSON(400, "Invalid Data")
		return
	}

	data, err := h.service.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) AuthUser(ctx *gin.Context) {
	user := domain.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.service.AuthUser(ctx, &user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

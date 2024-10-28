package v1

import (
	"anon-chat/internal/domain"
	"anon-chat/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authUsecase usecase.AuthUsecase
}

func newLoginRoutes(h *gin.RouterGroup, au *usecase.AuthUsecase) {
	r := &AuthRoutes{*au}
	{
		h.POST("/sign-up", r.SignUp)
		h.POST("/login", r.Login)
	}
}

func (r *AuthRoutes) SignUp(c *gin.Context) {
	var req domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, INVALID_REQUEST_BODY_MESSAGE)
		return
	}

	response, err := r.authUsecase.SignUp(c.Request.Context(), req)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *AuthRoutes) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, INVALID_REQUEST_BODY_MESSAGE)
		return
	}

	response, err := r.authUsecase.Login(c.Request.Context(), req)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

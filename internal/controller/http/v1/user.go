package v1

import (
	"context"
	"echo-box/internal/domain"
	"echo-box/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	UserUsecase   usecase.UserUsecase
	FriendUsecase usecase.FriendUsecase
}

func newUserRoutes(h *gin.RouterGroup, m *Middleware, uu *usecase.UserUsecase, fu *usecase.FriendUsecase) {
	r := &UserRoutes{*uu, *fu}
	{
		h.GET("/users", m.VerifyToken, r.GetUser)
		h.GET("/users/friends", m.VerifyToken, r.GetFriends)
	}
}

func (r *UserRoutes) GetUser(c *gin.Context) {
	user, ok := c.Get(string(userContextKey))
	if !ok {
		errorResponse(c, http.StatusBadRequest, "no user found for the given token")
	}
	response, err := r.UserUsecase.GetUser(context.WithValue(c, userContextKey, user))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *UserRoutes) GetFriends(c *gin.Context) {
	user, ok := c.Get(string(userContextKey))
	if !ok {
		errorResponse(c, http.StatusBadRequest, "no user found for the given token")
	}
	response, err := r.FriendUsecase.GetFriendsByUsername(user.(*domain.User).Username)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

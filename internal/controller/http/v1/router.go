package v1

import (
	"echo-box/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, m *Middleware, au *usecase.AuthUsecase, uu *usecase.UserUsecase,
	fu *usecase.FriendUsecase, wsM *WSManager, ex *Explorer) {
	handler.Use(m.CORSMiddleware)
	h := handler.Group("/echo-box/api/v1")
	h.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "service is up!!!")
	})
	{
		newLoginRoutes(h, au)
		newUserRoutes(h, m, uu, fu)
		newChatRoutes(h, m, wsM, uu, fu, ex)
	}
}

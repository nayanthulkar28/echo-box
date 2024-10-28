package v1

import (
	"anon-chat/internal/domain"
	"anon-chat/internal/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatRoutes struct {
	wsManager     *WSManager
	userUsecase   *usecase.UserUsecase
	friendUsecase *usecase.FriendUsecase
	explorer      *Explorer
}

func newChatRoutes(h *gin.RouterGroup, m *Middleware, wsM *WSManager, uu *usecase.UserUsecase,
	fu *usecase.FriendUsecase, ex *Explorer) {
	r := &ChatRoutes{wsManager: wsM, userUsecase: uu, friendUsecase: fu, explorer: ex}
	{
		h.GET("ws/chat", m.VerifyToken, r.ServerChatWS)
	}
}

func (r *ChatRoutes) ServerChatWS(c *gin.Context) {
	token := c.Request.URL.Query().Get("token")
	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("err", err)
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_user, exist := c.Get(string(userContextKey))
	if !exist {
		errorResponse(c, http.StatusBadRequest, "no user found for the given token")
		return
	}
	user := _user.(*domain.User)
	userData := domain.UserData{
		Username: user.Username,
	}

	client := NewWsClient(conn, r.friendUsecase)
	r.wsManager.addClient(userData, client)
	go client.readMessages(userData, r.wsManager, token, r.explorer)
}

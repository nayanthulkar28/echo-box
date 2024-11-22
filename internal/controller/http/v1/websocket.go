package v1

import (
	"echo-box/internal/domain"
	"echo-box/internal/usecase"
	"echo-box/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	ws "github.com/gorilla/websocket"
)

var (
	websocketUpgrader = ws.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WSClientList map[domain.UserData][]*WSClient

type WSClient struct {
	conn          *ws.Conn
	friendUsecase *usecase.FriendUsecase
}

func NewWsClient(conn *ws.Conn, fu *usecase.FriendUsecase) *WSClient {
	return &WSClient{
		conn:          conn,
		friendUsecase: fu,
	}
}

type WSManager struct {
	clients WSClientList
	sync.RWMutex
}

func NewWsManager() *WSManager {
	return &WSManager{
		clients: make(WSClientList),
	}
}

func (m *WSManager) addClient(user domain.UserData, client *WSClient) {
	m.Lock()
	defer m.Unlock()
	m.clients[user] = append(m.clients[user], client)
}

func (m *WSManager) removeClient(user domain.UserData, client *WSClient) {
	m.Lock()
	defer m.Unlock()
	if clientList, ok := m.clients[user]; ok {
		var index int
		for i := range clientList {
			if clientList[i] == client {
				index = i
				break
			}
		}
		clientList[index] = clientList[len(clientList)-1]
		m.clients[user] = clientList[:len(clientList)-1]
		if len(m.clients[user]) == 0 {
			delete(m.clients, user)
		}
		client.conn.Close()
	}
}

func (c *WSClient) readMessages(user domain.UserData, m *WSManager, tokenString string, explorer *Explorer) {
	defer func() {
		fmt.Printf("ending connection for user %s\n", user.Username)
		m.removeClient(user, c)
		explorer.RemoveExplorer(user)
	}()
	for {
		var payloadReq domain.ChatPayloadRequest
		var payloadRes domain.ChatPayloadResponse

		_, payloadReqBytes, err := c.conn.ReadMessage()
		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				fmt.Printf("error reading ws connection message %v\n", err)
			}
			break
		}

		token, err := pkg.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			if err = c.conn.WriteMessage(ws.CloseMessage, ws.FormatCloseMessage(http.StatusUnauthorized, "token expired")); err == nil {
				fmt.Printf("token is expired or invalid, closing connection...\n")
			}
			break
		}

		err = json.Unmarshal(payloadReqBytes, &payloadReq)
		if err != nil {
			fmt.Printf("error parsing payload %v", err)
			continue
		}

		switch payloadReq.Type {
		// handles message sharing, sending friend request and accepting friend request
		case domain.SEND_MESSAGE_TYPE, domain.SEND_FRIEND_REQUEST, domain.ACCEPT_FRIEND_REQUEST:
			payloadRes.Type = payloadReq.Type
			payloadRes.From.User = user
			payloadRes.Message = payloadReq.Message
			payloadResBytes, err := json.Marshal(payloadRes)
			if err != nil {
				fmt.Printf("error parsing payload %v", err)
				continue
			}

			if payloadReq.Type == domain.ACCEPT_FRIEND_REQUEST {
				err := c.friendUsecase.MakeFriends(user, payloadReq.To.User)
				if err != nil {
					fmt.Println("error making friends")
					continue
				}
				if toConnList, exist := m.clients[user]; exist {
					for _, toConn := range toConnList {
						addFriendPayload := domain.ChatPayloadResponse{
							From: domain.UserResponse{
								User: payloadReq.To.User,
							},
							Type:    payloadReq.Type,
							Message: payloadReq.Message,
						}
						addFriendPayloadBytes, err := json.Marshal(addFriendPayload)
						if err != nil {
							fmt.Printf("error parsing payload %v", err)
							continue
						}
						if err = toConn.conn.WriteMessage(ws.TextMessage, addFriendPayloadBytes); err != nil {
							fmt.Printf("error sending message %v", err)
						}
					}
					fmt.Println("message send")
				}
			}

			if toConnList, exist := m.clients[payloadReq.To.User]; exist {
				for _, toConn := range toConnList {
					if err = toConn.conn.WriteMessage(ws.TextMessage, payloadResBytes); err != nil {
						fmt.Printf("error sending message %v", err)
					}
				}
				fmt.Println("message send")
			} else {
				fmt.Printf("user %s is offline\n", payloadReq.To.User.Username)
			}
		case domain.ADD_EXPLORER_TYPE:
			explorer.AddExplorer(user)
			fmt.Println("explorer added")
		case domain.REMOVE_EXPLORER_TYPE:
			explorer.RemoveExplorer(user)
			fmt.Println("explorer removed")
		}
	}
}

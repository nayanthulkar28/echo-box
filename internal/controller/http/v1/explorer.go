package v1

import (
	"anon-chat/internal/domain"
	"anon-chat/internal/usecase"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/slices"

	ws "github.com/gorilla/websocket"
)

type Explorer struct {
	FriendUsecase *usecase.FriendUsecase
	List          chan domain.UserData
	Status        map[domain.UserData]bool
	Manager       *WSManager
	sync.RWMutex
}

func NewExplorer(m *WSManager, fu *usecase.FriendUsecase) *Explorer {
	return &Explorer{
		FriendUsecase: fu,
		List:          make(chan domain.UserData, 10),
		Status:        make(map[domain.UserData]bool),
		Manager:       m,
	}
}

func (e *Explorer) RunExplorer() {
	for {
		var user1, user2 domain.UserData
		var open bool
		for {
			user1, open = <-e.List
			if !open {
				return
			}
			if _, ok := e.Status[user1]; ok {
				e.RemoveExplorer(user1)
				break
			}
		}

		user1FriendList, err := e.FriendUsecase.GetFriendsByUsername(user1.Username)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for {
			user2, open = <-e.List
			if !open {
				return
			}
			exist := slices.Contains(user1FriendList, user2)
			if exist {
				e.List <- user2
				time.Sleep(1 * time.Second)
				continue
			}
			if _, ok := e.Status[user2]; ok {
				e.RemoveExplorer(user2)
				break
			}
		}

		var payloadRes1 domain.ChatPayloadResponse
		var payloadRes2 domain.ChatPayloadResponse

		payloadRes1.From = domain.UserResponse{User: user2}
		payloadRes1.Type = domain.ADD_EXPLORER_TYPE
		payloadRes2.From = domain.UserResponse{User: user1}
		payloadRes2.Type = domain.ADD_EXPLORER_TYPE

		payloadRes1Bytes, _ := json.Marshal(payloadRes1)
		payloadRes2Bytes, _ := json.Marshal(payloadRes2)

		var connUser1List, connUser2List []*WSClient
		var exist bool

		if connUser1List, exist = e.Manager.clients[user1]; !exist {
			fmt.Println("user1 disconnected")
			continue
		}
		if connUser2List, exist = e.Manager.clients[user2]; !exist {
			fmt.Println("user2 disconnected")
			continue
		}

		for _, connUser1 := range connUser1List {
			if err := connUser1.conn.WriteMessage(ws.TextMessage, payloadRes1Bytes); err != nil {
				fmt.Println("error sending explorer response")
				continue
			}
			fmt.Println("explorer send for", user1)
		}
		for _, connUser2 := range connUser2List {
			if err := connUser2.conn.WriteMessage(ws.TextMessage, payloadRes2Bytes); err != nil {
				fmt.Println("error sending explorer response")
				continue
			}
			fmt.Println("explorer send for", user2)
		}
	}
}

func (e *Explorer) AddExplorer(user domain.UserData) {
	e.List <- user
	e.Lock()
	defer e.Unlock()
	e.Status[user] = true
}

func (e *Explorer) RemoveExplorer(user domain.UserData) {
	e.Lock()
	defer e.Unlock()
	delete(e.Status, user)
}

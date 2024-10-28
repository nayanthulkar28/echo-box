package domain

const (
	SEND_MESSAGE_TYPE     = 1
	ADD_EXPLORER_TYPE     = 2
	REMOVE_EXPLORER_TYPE  = 3
	SEND_FRIEND_REQUEST   = 4
	ACCEPT_FRIEND_REQUEST = 5
)

type ChatPayloadRequest struct {
	To      UserResponse `json:"to"`
	Type    int          `json:"type"`
	Message string       `json:"message"`
}

type ChatPayloadResponse struct {
	From    UserResponse `json:"from"`
	Type    int          `json:"type"`
	Message string       `json:"message"`
}

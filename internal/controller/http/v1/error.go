package v1

import "github.com/gin-gonic/gin"

type response struct {
	Error string `json:"error"`
}

const (
	INVALID_REQUEST_BODY_MESSAGE = "invalid request body"
)

func errorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, response{message})
}

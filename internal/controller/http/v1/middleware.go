package v1

import (
	"echo-box/internal/repo"
	"echo-box/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type ContextKey string

const userContextKey = ContextKey("user")

type Middleware struct {
	UserRepo *repo.UserRepo
}

func NewMiddleware(ar *repo.UserRepo) *Middleware {
	return &Middleware{
		UserRepo: ar,
	}
}

func (m *Middleware) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Accept,X-Requested-With,Content-Type,Authorization")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func (m *Middleware) VerifyToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		authHeader = c.Request.URL.Query().Get("token")
	}
	token, err := pkg.VerifyToken(authHeader)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	payload := token.Claims.(jwt.MapClaims)
	username := payload["username"].(string)
	user, err := m.UserRepo.GetUserByUsername(c, username)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(string(userContextKey), user)
	c.Next()
}

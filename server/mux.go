package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lisuiheng/fsm"
	"net/http"
)

// MuxOpts are the options for the router.  Mostly related to auth.
type MuxOpts struct {
	Logger    fsm.Logger
	Signature string // UseAuth turns on Github OAuth and JWT
}

// NewMux attaches all the route handlers; handler returned servers chronograf.
func NewMux(opts MuxOpts, service Service) http.Handler {
	router := gin.Default()
	authorized := router.Group("/")

	router.POST("/api/login", service.Login)

	authorized.Use(service.auth)
	{
		authorized.GET("/api/users", service.User)
		authorized.POST("/api/users", service.NewUser)
		authorized.GET("/api/users/:id", service.UserID)
		authorized.PUT("/api/users/:id", service.UpdateUser)
		authorized.DELETE("/api/users/:id", service.RemoveUser)
	}

	return router
}

func (s *Service) unauthorizedError(c *gin.Context, err error) {
	s.errorw(c, http.StatusUnauthorized, fmt.Sprintf("Unauthorized error: %v", err))
}

func (s *Service) invalidJSON(c *gin.Context) {
	s.errorw(c, http.StatusBadRequest, "Unparsable JSON")
}

func (s *Service) unknownError(c *gin.Context, err error) {
	s.errorw(c, http.StatusInternalServerError, fmt.Sprintf("Unknown error: %v", err))
}

func (s *Service) forbiddenError(c *gin.Context) {
	s.errorw(c, http.StatusForbidden, fmt.Sprintf("username or password invalid"))
}

func (s *Service) notFound(c *gin.Context, id interface{}) {
	s.errorw(c, http.StatusNotFound, fmt.Sprintf("ID %v not found", id))
}

func (s *Service) errorw(c *gin.Context, code int, msg string) {
	s.Logger.Errorf("server error code:%d, msg:%s", code, msg)
	c.AbortWithStatusJSON(code, gin.H{"message": msg})
}

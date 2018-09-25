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

	authorized.Use(Middleware())

	authorized.GET("/api/file", service.upload)

	return router
}

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Query("id")
		fmt.Println("Middleware", id)
		c.JSON(http.StatusOK, "Middleware")
		c.Next()
	}

}

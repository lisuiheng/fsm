package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) upload(c *gin.Context) {
	id := c.Query("id")

	fmt.Println("upload", id, c.Writer.Status())
	c.Writer.Hijack()
	c.JSON(http.StatusOK, "upload")
}

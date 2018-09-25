package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/api/proxy", func(c *gin.Context) {
		c.JSON(200, "2")
	})
	r.Run(":8002") // listen and serve on 0.0.0.0:8080
}

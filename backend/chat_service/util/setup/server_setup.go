package setup

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ServerSetup() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello, world!")
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.Run("0.0.0.0:8003")
}

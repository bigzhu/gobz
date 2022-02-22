package main

import (
	"github.com/bigzhu/gobz/websocketbz"
	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	m.HandleMessage(websocketbz.DoRegister)

	r.Run(":5000")
}

package main

import (
	"log"

	"github.com/bigzhu/gobz/test"
	"github.com/bigzhu/gobz/websocketbz"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//SetupPlay 运维平台
func SetupPlay() *gin.Engine {
	r := gin.New()
	playGroup := r.Group("/playAPI")
	playGroup.GET("Google", test.Google)
	r.GET("/ws", func(c *gin.Context) {
		websocketbz.M.HandleRequest(c.Writer, c.Request)
	})
	return r
}
func main() {
	log.Println("启动 gobz")
	r := SetupPlay()
	err := r.Run(":3003")
	if err != nil {
		panic(err)
	}
}

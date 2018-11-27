package main

import (
	"log"

	_ "github.com/bigzhu/gobz/apibz"
	_ "github.com/bigzhu/gobz/confbz"
	_ "github.com/bigzhu/gobz/functionbz"
	_ "github.com/bigzhu/gobz/ginbz"
	_ "github.com/bigzhu/gobz/httpbz"
	_ "github.com/bigzhu/gobz/jsonbz"
	_ "github.com/bigzhu/gobz/modelsbz"
	_ "github.com/bigzhu/gobz/rsabz"
	_ "github.com/bigzhu/gobz/services/callbackbz"
	_ "github.com/bigzhu/gobz/services/mailbz"
	_ "github.com/bigzhu/gobz/services/oauthbz"
	_ "github.com/bigzhu/gobz/services/smsbz"
	_ "github.com/bigzhu/gobz/telegram"
	"github.com/bigzhu/gobz/test"
	"github.com/bigzhu/gobz/websocketbz"
	_ "github.com/bigzhu/gobz/websocketbz"
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

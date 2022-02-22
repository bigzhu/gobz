package main

import (
	"log"

	_ "gobz/apibz"
	_ "gobz/confbz"
	_ "gobz/functionbz"
	_ "gobz/ginbz"
	_ "gobz/httpbz"
	_ "gobz/jsonbz"
	_ "gobz/modelsbz"
	_ "gobz/rsabz"
	_ "gobz/services/callbackbz"
	_ "gobz/services/mailbz"
	_ "gobz/services/oauthbz"
	_ "gobz/services/smsbz"
	_ "gobz/services/tokenize"
	_ "gobz/telegram"
	"gobz/test"
	"gobz/websocketbz"
	_ "gobz/websocketbz"

	"github.com/gin-gonic/gin"
	_ "github.com/ugorji/go/codec"
	_ "gorm.io/driver/postgres"
)

//SetupPlay 运维平台
func SetupPlay() *gin.Engine {
	r := gin.New()
	api := r.Group("/api")
	api.GET("google", test.Google)
	api.GET("github", test.Github)
	r.GET("/ws", func(c *gin.Context) {
		websocketbz.M.HandleRequest(c.Writer, c.Request)
	})
	return r
}
func main() {
	// 显示文件路径
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("启动 gobz")
	r := SetupPlay()
	err := r.Run(":3003")
	if err != nil {
		panic(err)
	}
}

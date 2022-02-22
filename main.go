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
	_ "github.com/bigzhu/gobz/services/tokenize"
	_ "github.com/bigzhu/gobz/telegram"
	"github.com/bigzhu/gobz/test"
	"github.com/bigzhu/gobz/websocketbz"
	_ "github.com/bigzhu/gobz/websocketbz"

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
	log.Println("启动 github.com/bigzhu/gobz")
	r := SetupPlay()
	err := r.Run(":3003")
	if err != nil {
		panic(err)
	}
}

package oauthbz

import (
	"net/http"

	"github.com/bigzhu/gobz/apibz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// GetAccessToken 获取到AccessToken
// c gin 的环境变量
// conf oauth2 的相关配置
func GetAccessToken(c *gin.Context, conf *oauth2.Config) (accessToken string, err error) {

	state := c.Query("state")
	// 如果没有传递 state 参数过来, 重定向到获取 state 页面(授权)
	if state == "" {
		url := conf.AuthCodeURL("state")
		c.Redirect(http.StatusFound, url)
		return
	}
	// 如果有传 code, 用 code 换取 token
	code := c.Query("code")
	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apibz.NewE(err))
		return
	}
	accessToken = token.AccessToken
	return
}

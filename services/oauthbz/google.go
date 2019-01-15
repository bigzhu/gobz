package oauthbz

// oauth 登录相关东西. 依赖 gin

import (
	"encoding/json"
	"log"

	"github.com/bigzhu/gobz/httpbz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google google 的类型
const Google = "google"

// OauthGoogle oauth2
func OauthGoogle(c *gin.Context) (googleUserInfo OauthInfo, err error) {
	googleOauthConf := &oauth2.Config{
		ClientID:     oauthConf.Google.ClientID,
		ClientSecret: oauthConf.Google.ClientSecret,
		RedirectURL:  oauthConf.Google.RedirectURL,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	accessToken, err := GetAccessToken(c, googleOauthConf)
	data, _, err := httpbz.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+accessToken, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &googleUserInfo)
	googleUserInfo.Type = Google
	d, _ := json.Marshal(googleUserInfo)
	log.Println(string(d))
	return
}

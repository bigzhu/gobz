package oauthbz

// oauth 登录相关东西. 依赖 gin

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bigzhu/gobz/httpbz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google google 的类型
const Google = "google"

// OauthGoogle oauth2
func OauthGoogle(c *gin.Context, redirectURL string, clientID string, clientSecret string) (googleUserInfo OauthInfo, err error) {
	// c := outc.(*gin.Context)
	var googleConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	state := c.Query("state")
	if state == "" {
		url := googleConf.AuthCodeURL("state")
		c.Redirect(http.StatusFound, url)
	}
	code := c.Query("code")
	token, err := googleConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ERROR: err.Error()})
		return
	}

	data, err := httpbz.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &googleUserInfo)
	googleUserInfo.Type = Google
	log.Println(googleUserInfo)
	return
}

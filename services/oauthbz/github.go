package oauthbz

// oauth 登录相关东西. 依赖 gin

import (
	"log"

	"github.com/bigzhu/gobz/httpbz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GitHub github 的类型
const GitHub = "github"

// OauthGithub oauth2
func OauthGithub(c *gin.Context) (googleUserInfo OauthInfo, err error) {
	githubOauthConf := &oauth2.Config{
		ClientID:     oauthConf.GitHub.ClientID,
		ClientSecret: oauthConf.GitHub.ClientSecret,
		RedirectURL:  oauthConf.GitHub.RedirectURL,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
	accessToken, err := GetAccessToken(c, githubOauthConf)
	data, _, err := httpbz.Get("https://api.github.com/user?access_token="+accessToken, nil)
	if err != nil {
		return
	}
	log.Println(data)

	//err = json.Unmarshal([]byte(data), &googleUserInfo)
	//googleUserInfo.Type = Google
	//d, _ := json.Marshal(googleUserInfo)
	//log.Println(string(d))
	return
}

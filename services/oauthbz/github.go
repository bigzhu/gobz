package oauthbz

// oauth 登录相关东西. 依赖 gin

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gobz/apibz"
	"gobz/httpbz"
	"gobz/modelsbz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Github github 的类型
const Github = "github"

// OauthGithub oauth2
func OauthGithub(c *gin.Context) (oauthInfo modelsbz.OauthInfo, err error) {
	githubOauthConf := &oauth2.Config{
		ClientID:     oauthConf.GitHub.ClientID,
		ClientSecret: oauthConf.GitHub.ClientSecret,
		RedirectURL:  oauthConf.GitHub.RedirectURL,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
	accessToken, err := GetAccessToken(c, githubOauthConf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apibz.NewE(err))
		return
	}
	if accessToken == "" { // 说明还在前奏
		return
	}
	data, _, err := httpbz.Get("https://api.github.com/user?access_token="+accessToken, nil)
	if err != nil {
		return
	}
	githubOauthInfo := GithubOauthInfo{}
	err = json.Unmarshal([]byte(data), &githubOauthInfo)
	oauthInfo.Type = Github
	oauthInfo.Email = githubOauthInfo.Email
	oauthInfo.OutID = strconv.Itoa(githubOauthInfo.ID)
	oauthInfo.Name = githubOauthInfo.Name
	oauthInfo.AvatarURL = githubOauthInfo.AvatarURL
	oauthInfo.Link = githubOauthInfo.HTMLURL
	d, _ := json.Marshal(oauthInfo)
	log.Println(string(d))
	return
}

// OauthGithubSave oauth 登录并且保存
func OauthGithubSave(c *gin.Context) (oauthInfo modelsbz.OauthInfo, err error) {
	oauthInfo, err = OauthGithub(c)
	if err != nil {
		return
	}
	err = SaveOrGet(&oauthInfo)
	return
}

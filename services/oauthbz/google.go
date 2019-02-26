package oauthbz

// oauth 登录相关东西. 依赖 gin

import (
	"encoding/json"
	"log"

	"github.com/bigzhu/gobz/httpbz"
	"github.com/bigzhu/gobz/modelsbz"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google google 的类型
const Google = "google"

// OauthGoogle oauth2
func OauthGoogle(c *gin.Context) (oauthInfo modelsbz.OauthInfo, err error) {
	googleOauthConf := &oauth2.Config{
		ClientID:     oauthConf.Google.ClientID,
		ClientSecret: oauthConf.Google.ClientSecret,
		RedirectURL:  oauthConf.Google.RedirectURL,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	accessToken, err := GetAccessToken(c, googleOauthConf)
	if err != nil {
		return
	}
	if accessToken == "" { // 说明还在前奏
		return
	}
	data, _, err := httpbz.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+accessToken, nil)
	if err != nil {
		return
	}
	googleOauthInfo := GoogleOauthInfo{}
	err = json.Unmarshal([]byte(data), &googleOauthInfo)
	if err != nil {
		return
	}
	oauthInfo.Type = Google
	oauthInfo.Email = googleOauthInfo.Email
	oauthInfo.OutID = googleOauthInfo.ID
	oauthInfo.Name = googleOauthInfo.Name
	oauthInfo.AvatarURL = googleOauthInfo.Picture
	oauthInfo.Link = googleOauthInfo.Link
	d, _ := json.Marshal(oauthInfo)
	log.Println(string(d))
	return
}

// OauthGoogleSave oauth 登录并且保存
func OauthGoogleSave(c *gin.Context) (oauthInfo modelsbz.OauthInfo, err error) {
	oauthInfo, err = OauthGoogle(c)
	if err != nil {
		return
	}
	err = SaveOrGet(&oauthInfo)
	return
}

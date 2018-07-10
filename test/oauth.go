package test

import (
	"github.com/bigzhu/gobz/confbz"
	"github.com/bigzhu/gobz/ginbz"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	oauthConf := confbz.GetOauthConf()
	_, _ = ginbz.OauthGoogle(c, oauthConf.RedirectURL, oauthConf.ClientID, oauthConf.ClientSecret)
}

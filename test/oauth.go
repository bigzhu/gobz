package test

import (
	"github.com/bigzhu/gobz/confbz"
	"github.com/bigzhu/gobz/oauthbz"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	oauthConf := confbz.GetOauthConf()
	_, _ = oauthbz.OauthGoogle(c, oauthConf.RedirectURL, oauthConf.ClientID, oauthConf.ClientSecret)
}

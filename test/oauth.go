package test

import (
	"github.com/bigzhu/gobz/confbz"
	"github.com/bigzhu/gobz/services/oauthbz"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	oauthConf := confbz.GetOauthConf()
	google := oauthConf.Google
	_, _ = oauthbz.OauthGoogle(c, google.RedirectURL, google.ClientID, google.ClientSecret)
}

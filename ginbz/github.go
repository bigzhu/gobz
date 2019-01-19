package ginbz

import (
	"net/http"

	"github.com/bigzhu/gobz/apibz"
	"github.com/bigzhu/gobz/services/oauthbz"
	"github.com/gin-gonic/gin"
)

// Github oauth2
func Github(c *gin.Context) {
	oauthInfo, err := oauthbz.OauthGithubSave(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apibz.NewE(err))
		return
	}
	err = saveOauthSession(c, oauthInfo.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apibz.NewE(err))
		return
	}
	c.Redirect(http.StatusFound, "/")
}

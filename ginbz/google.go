package ginbz

import (
	"net/http"

	"github.com/bigzhu/gobz/apibz"
	"github.com/bigzhu/gobz/services/oauthbz"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Google oauth2
func Google(c *gin.Context) {
	oauthInfo, err := oauthbz.OauthGoogleSave(c)
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

func saveOauthSession(c *gin.Context, id int) (err error) {
	session := sessions.Default(c)
	session.Set("user", id)
	err = session.Save()
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

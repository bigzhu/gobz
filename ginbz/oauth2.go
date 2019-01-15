package ginbz

import (
	"net/http"

	"github.com/bigzhu/gobz/apibz"
	"github.com/bigzhu/gobz/modelsbz"
	"github.com/bigzhu/gobz/services/oauthbz"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	oauthInfo, err := oauthbz.OauthGoogle(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apibz.NewE(err))
		return
	}
	err = modelsbz.DB.Where("email=?", oauthInfo.Email).FirstOrCreate(&oauthInfo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apibz.NewE(err))
		return
	}
	session := sessions.Default(c)
	session.Set("user", oauthInfo.ID)
	err = session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apibz.NewE(err))
		return
	}
	c.Redirect(http.StatusFound, "/")
}

// OauthInfo 获取用户信息
func OauthInfo(c *gin.Context) {
	userID := getUserID(c)
	oauth := oauthbz.OauthInfo{}
	modelsbz.DB.
		Where("id=?", userID).
		Find(&oauth)
	c.JSON(http.StatusOK, oauth)
}

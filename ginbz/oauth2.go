package ginbz

import (
	"net/http"

	"github.com/bigzhu/gobz/confbz"
	"github.com/bigzhu/gobz/modelsbz"
	"github.com/bigzhu/gobz/services/oauthbz"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	oauthConf := confbz.GetOauthConf()
	oauthInfo, err := oauthbz.OauthGoogle(c, oauthConf.RedirectURL, oauthConf.ClientID, oauthConf.ClientSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ERROR: err.Error()})
		return
	}
	oauthInfo.Type = "google"
	err = modelsbz.DB.Where("email=?", oauthInfo.Email).FirstOrCreate(&oauthInfo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ERROR: err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("user", oauthInfo.OID)
	err = session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ERROR: err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

// OauthInfo 获取用户信息
func OauthInfo(c *gin.Context) {
	userID := getUserID(c)
	oauth := oauthbz.OauthInfo{OID: userID}
	modelsbz.DB.
		Where("o_id=?", userID).
		Find(&oauth)
	c.JSON(http.StatusOK, oauth)
}

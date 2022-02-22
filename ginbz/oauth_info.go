package ginbz

import (
	"net/http"

	"github.com/bigzhu/gobz/modelsbz"
	"github.com/gin-gonic/gin"
)

// OauthInfo 获取用户信息
func OauthInfo(c *gin.Context) {
	userID := getUserID(c)
	oauth := modelsbz.OauthInfo{}
	modelsbz.DB.
		Where("id=?", userID).
		Find(&oauth)
	c.JSON(http.StatusOK, oauth)
}

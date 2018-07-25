package ginbz

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	userID := session.Get("user")
	if userID != nil {
		return userID.(uint)
	}
	return 0
}

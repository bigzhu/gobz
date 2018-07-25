package ginbz

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Logout 退出
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user", "")
	err := session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ERROR: err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

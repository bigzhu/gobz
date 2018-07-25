package ginbz

import (
	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) uint {
	userID, exists := c.Get("userID")
	if !exists {
		// c.JSON(http.StatusUnauthorized, gin.H{ERROR: "未正确取到登录的用户ID"})
		return 0
	}
	return userID.(uint)
}

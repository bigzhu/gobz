package ginbz

import (
	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) uint {
	// 在 middleware 核对时候应该设置了 userID
	userID, exists := c.Get("userID")
	if exists {
		return userID.(uint)
	}
	return 0
}

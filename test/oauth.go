package test

import (
	"github.com/bigzhu/gobz/services/oauthbz"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	_, _ = oauthbz.OauthGoogle(c)
}

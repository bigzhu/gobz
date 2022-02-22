package test

import (
	"gobz/services/oauthbz"
	"github.com/gin-gonic/gin"
)

// Google oauth2
func Google(c *gin.Context) {
	_, _ = oauthbz.OauthGoogle(c)
}

// Github oauth2
func Github(c *gin.Context) {
	_, _ = oauthbz.OauthGithub(c)
}

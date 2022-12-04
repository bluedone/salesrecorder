package service

import (
	b64 "encoding/base64"
	"net/http"
	"strings"

	config "sr-server/config"

	"github.com/gin-gonic/gin"
)

func GetHeaderAuth(c *gin.Context) int {
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	return GetUserFromToken(token)
}

func AdminAuth(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Basic ")
	str := []string{config.AdminUser, config.AdminPassword}
	if b64.StdEncoding.EncodeToString([]byte(strings.Join(str, ":"))) != token {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
		return
	}
}

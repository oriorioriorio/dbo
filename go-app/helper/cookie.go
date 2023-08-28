package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool) {
	c.SetSameSite(http.SameSiteNoneMode)
	// set secure='true' in prod
	c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

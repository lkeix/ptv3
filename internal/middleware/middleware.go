package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Authorize is middleware of authorize
func Authorize(ctx *gin.Context) {
	session := sessions.Default(ctx)
	log := session.Get("login")
	if log == nil || !log.(bool) {
		ctx.Redirect(303, "/signerror")
		// ctx.Abort()
		return
	}
}

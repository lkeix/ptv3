package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// TabSelect select tab data.
func TabSelect(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	tab := req.PostFormValue("tab")
	session := sessions.Default(ctx)
	session.Set("tab", tab)
	session.Save()
}

func tabProcess(tab string) {

}

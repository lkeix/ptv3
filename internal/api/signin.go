package api

import (
	"ptv3/internal/dblib"
	utils "ptv3/internal/util"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
)

// Login is login api point, Login is called post login request.
func Login(ctx *gin.Context) {
	id := ctx.PostForm("id")
	password := ctx.PostForm("password")
	hashed := dblib.ReadDoctorAuthorizeInfo(id)
	certification := utils.ComparePasswd(password, hashed)
	if certification {
		session := sessions.Default(ctx)
		session.Set("login", true)
		session.Set("doctorID", id)
		session.Set("step", int64(1))
		session.Save()
		ctx.Redirect(303, "/ptv3/home")
	} else {
		ctx.Redirect(303, "/ptv3/signerror")
	}
}

// Signup is Signup api point, Signup is called post signup request.
func Signup(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var isMaster bool
	id := ctx.PostForm("id")
	password := ctx.PostForm("password")
	institute := ctx.PostForm("institute")
	masterKey := ctx.PostForm("masterKey")
	defaultKey := "xbkYF6UbDu7LW8Xf"
	if masterKey == defaultKey {
		isMaster = true
	}
	instituteID := dblib.ReadInstitutionID(institute)
	if instituteID <= 0 {
		ctx.Redirect(303, "/ptv3/signerror")
	}
	hashed := utils.CryptPasswd(password)
	if !dblib.InsertNewDoctor(id, hashed, isMaster, instituteID) {
		ctx.Redirect(303, "/ptv3/signerror")
	} else {
		session.Set("login", true)
		session.Set("doctorID", id)
		session.Set("step", int64(1))
		session.Save()
		ctx.Redirect(303, "/ptv3/home")
	}
}

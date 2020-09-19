package route

import (
	"net/http"
	"ptv3/internal/api"
	"ptv3/internal/middleware"
	"strconv"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Menu is nav bar menu data.
type Menu struct {
	Name string
	Link string
}

// Route define Application router
func Route() *gin.Engine {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	templates := multitemplate.New()
	templates.AddFromFiles(
		"home",
		"templates/home.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
		"templates/list/view.tmpl",
		"templates/list/delete.tmpl",
	)
	templates.AddFromFiles(
		"signin",
		"templates/signin.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
	)
	templates.AddFromFiles(
		"logger",
		"templates/logger/controller.tmpl",
		"templates/logger/form.tmpl",
		"templates/logger/explain.tmpl",
		"templates/logger/work.tmpl",
		"templates/logger/check.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
	)
	templates.AddFromFiles(
		"signerror",
		"templates/signerror.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
	)
	templates.AddFromFiles(
		"delete",
		"templates/delete.tmpl",
		"templates/list/delete.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
	)
	r.HTMLRender = templates
	r.Static("/js", "public/js")
	r.Static("/css", "public/css")

	r.GET("/signin", signin)
	r.GET("/signerror", signerror)
	r.POST("/api/login", api.Login)
	r.POST("/api/signup", api.Signup)
	requireAuth := r.Group("")
	requireAuth.Use(middleware.Authorize)
	{
		requireAuth.GET("/home", home)
		requireAuth.GET("/logger", logger)
		requireAuth.GET("/logout", logout)
		requireAuth.GET("/delete", delete)
	}

	requireAuthAPI := r.Group("api")
	{
		requireAuthAPI.POST("/newPatientStart", api.NewPatientStart)
		requireAuthAPI.POST("/nextstep", api.NextStep)
		requireAuthAPI.POST("/storeLog", api.StoreLog)
		requireAuthAPI.POST("/readPatientData", api.ReadPatientData)
		requireAuthAPI.POST("/readPastPatient", api.ReadPastPatients)
		requireAuthAPI.POST("/pastPatientStart", api.PastPatientStart)
		requireAuthAPI.POST("/deletePatient", api.DeletePatient)
		requireAuthAPI.POST("/readAllPatients", api.GetAllPatientsData)
	}
	return r
}

func home(ctx *gin.Context) {
	var exist bool
	var page interface{}
	if page, exist = ctx.GetQuery("page"); !exist {
		page = 1
	} else {
		page, _ = strconv.Atoi(page.(string))
	}
	patientsData, page := api.GetPatientsData(ctx, page)
	ctx.HTML(http.StatusOK, "home", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "home",
		"Patients": patientsData,
		"Page":     page,
	})
}

func signin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "signin",
	})
}

func logger(ctx *gin.Context) {
	step := api.LoggerStep(ctx)
	ctx.HTML(http.StatusOK, "logger", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "logger",
		"step":     step,
	})
}

func signerror(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signerror", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "logger",
	})
}

func isLogin(ctx *gin.Context) []Menu {
	session := sessions.Default(ctx)
	var menu []Menu
	if session.Get("login") != nil {
		menu = []Menu{
			{
				Name: "ホーム",
				Link: "/home",
			},
			{
				Name: "ログ",
				Link: "/logger",
			},
			{
				Name: "欠損計算",
				Link: "/calcLoss",
			},
			{
				Name: "解析",
				Link: "/analysis",
			},
			{
				Name: "ログアウト",
				Link: "/logout",
			},
		}
	} else {
		menu = []Menu{
			{
				Name: "ログイン",
				Link: "/signin",
			},
		}
	}
	return menu
}

func delete(ctx *gin.Context) {
	var exist bool
	var page interface{}
	if page, exist = ctx.GetQuery("page"); !exist {
		page = 1
	} else {
		page, _ = strconv.Atoi(page.(string))
	}
	patientsData := api.GetAllPatients(ctx)
	ctx.HTML(http.StatusOK, "delete", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "delete",
		"Patients": patientsData,
		"Page":     page,
	})
}

func logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(303, "/signin")
}

package route

import (
	"net/http"
	"ptv3/internal/api"
	"ptv3/internal/dblib"
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
	options := sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   3600 * 24,
		Secure:   false,
		HttpOnly: false,
	}
	store.Options(options)
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
		"templates/error/signerror.tmpl",
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
	templates.AddFromFiles(
		"loss",
		"templates/loss/loss.tmpl",
		"templates/loss/list.tmpl",
		"templates/loss/result.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
	)
	templates.AddFromFiles(
		"notFound",
		"templates/error/notFound.tmpl",
		"templates/base/base.tmpl",
		"templates/base/header.tmpl",
	)
	templates.AddFromFiles(
		"analysis",
		"templates/analysis/controller.tmpl",
		"templates/analysis/tab.tmpl",
		"templates/analysis/recipe.tmpl",
		"templates/analysis/principle.tmpl",
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
	r.NoRoute(notFound)
	requireAuth := r.Group("")
	requireAuth.Use(middleware.Authorize)
	{
		requireAuth.GET("/home", home)
		requireAuth.GET("/logger", logger)
		requireAuth.GET("/logout", logout)
		requireAuth.GET("/delete", delete)
		requireAuth.GET("/loss", loss)
		requireAuth.POST("/lossCalc", lossCalc)
		requireAuth.GET("/analysis", analysis)
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
		requireAuthAPI.POST("/TabSelect", api.TabSelect)
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
				Link: "/loss",
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

func loss(ctx *gin.Context) {
	patientsData := api.GetAllPatients(ctx)
	ctx.HTML(http.StatusOK, "loss", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "loss",
		"view":     "listView",
		"Patients": patientsData,
	})
}

func lossCalc(ctx *gin.Context) {
	res := api.CalclateLoss(ctx)
	ctx.HTML(http.StatusOK, "loss", gin.H{
		"Menus":       isLogin(ctx),
		"Location":    "loss",
		"view":        "resultView",
		"Dates":       res.Dates,
		"ExplainLoss": res.ExplainLoss,
		"WorkLoss":    res.WorkLoss,
		"Rng":         res.Rng,
	})
}

func notFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "notFound", gin.H{
		"Menus":    isLogin(ctx),
		"Location": "notFound",
	})
}

func getTab(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	res := "recipe"
	tab := session.Get("tab")
	if tab != nil {
		res = tab.(string)
	}
	return res
}

func analysis(ctx *gin.Context) {
	tab := getTab(ctx)
	session := sessions.Default(ctx)
	doctorID := session.Get("doctorID").(string)
	recipeTab := "nav-link active"
	principleTab := "nav-link"
	if tab == "principle" {
		recipeTab = "nav-link"
		principleTab = "nav-link active"
	}
	patientsData := dblib.ReadAllPatientsLog(doctorID)
	ctx.HTML(http.StatusOK, "analysis", gin.H{
		"Menus":        isLogin(ctx),
		"Location":     "analysis",
		"recipeTab":    recipeTab,
		"principleTab": principleTab,
		"select":       tab,
		"patientsData": patientsData,
	})
}

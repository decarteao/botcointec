package routes

import (
	"botcointec/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistarLPRotas(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/login", middlewares.LoginAndSingupForm(), func(ctx *gin.Context) {
		flash, err := ctx.Cookie("flash")

		if err != nil {
			flash = ""
		} else {
			ctx.SetCookie("flash", "", -1, "", "", false, false)
		}

		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Logue no sistema",
			"flash": flash,
		})
	})

	r.GET("/signup", middlewares.LoginAndSingupForm(), func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Cadastre-se",
		})
	})
}

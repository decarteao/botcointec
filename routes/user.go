package routes

import (
	"botcointec/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistarUserRotas(r *gin.Engine) {
	user := r.Group("/dashboard")

	// middlewares
	user.Use(middlewares.AuthMiddleware("user"))

	user.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_index.html", gin.H{
			"title": "Seja bem-vindo",
		})
	})

	user.GET("/setup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_setup.html", gin.H{
			"title": "Setup do Bot",
		})
	})

	user.GET("/depositos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_depositos.html", gin.H{
			"title": "Depósitos",
		})
	})

	user.GET("/afiliados", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_afiliados.html", gin.H{
			"title": "Depósitos",
		})
	})

	user.GET("/geral", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_geral.html", gin.H{
			"title": "Configurações",
		})
	})

	user.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/login")
	})
}

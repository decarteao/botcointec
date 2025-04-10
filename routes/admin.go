package routes

import (
	"botcointec/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistarAdminRotas(r *gin.Engine) {
	admin := r.Group("/admin")

	// middlewares
	admin.Use(middlewares.AuthMiddleware("admin"))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_index.html", gin.H{
			"title": "Bem vindo"})
	})

	admin.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_users.html", gin.H{
			"title": "Gerencimento de usuários",
		})
	})

	admin.GET("/geral", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_geral.html", gin.H{
			"title": "Configurações",
		})
	})

	admin.GET("/sacar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_sacar.html", gin.H{
			"title": "Sacar Fundos",
		})
	})

	admin.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/login")
	})
}

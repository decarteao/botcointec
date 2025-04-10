package routes

import (
	"botcointec/config"
	"botcointec/controllers"
	"botcointec/models"
	"botcointec/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJSON struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Code       string `json:"code"` // codigo do google recaptcha
	RememberMe bool   `json:"remember_me"`
}
type RegisterJSON struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"` // codigo do google recaptcha
	Code     string `json:"code"`  // codigo do google recaptcha
}

func RegisterApi(r *gin.Engine, db *config.DataBase) {
	api := r.Group("/api")

	api.POST("/login", func(ctx *gin.Context) {
		var lj LoginJSON

		// passar o json do cliente pra variavel do servidor
		ctx.BindJSON(&lj)

		var user models.User

		// pesquisar no db
		if err := db.DB.Where("username = ?", lj.Username).First(&user).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"error":  "Credenciais incorrectas",
				// mesmas respostas para o password, para evitar enumeracao de usuarios
			})
		} else {
			// achou o usuario
			// verificar hash
			if utils.ValidarHash(lj.Password, user.Password) {
				// login efectuado com sucesso

				token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
					"user_id":    user.ID,
					"expiration": time.Now().Add(time.Hour * 168).Local().Unix(), // 7 dias
					"role":       user.Role,
					"remember":   lj.RememberMe,
				})

				tokenString, _ := token.SignedString(utils.JwtSecrect)

				var redirect_to = "/dashboard"
				if user.Role == "admin" {
					redirect_to = "/admin"
				}

				// setar os cookies
				ctx.SetCookie("token", tokenString, 604800, "/", "", false, true) // cookies validos ate 7 dias

				// enviar pro cliente
				ctx.JSON(http.StatusOK, gin.H{
					"status": true,
					// "token":       tokenString,
					"redirect_to": redirect_to,
				})
			} else {
				// erro na password
				ctx.JSON(http.StatusUnauthorized, "Credenciais incorrectas")
				// mesmas respostas para o password, para evitar enumeracao de usuarios
			}
		}
	})

	// registro
	api.POST("/register", func(ctx *gin.Context) {
		var rj RegisterJSON

		// passar o json do cliente pra variavel do servidor
		ctx.BindJSON(&rj)

		var user models.User

		// pesquisar no db
		if err := db.DB.Where("username = ? or email = ?", rj.Username, rj.Email).First(&user).Error; err != nil {
			// Cadastrar user
			controllers.AddUser(db, rj.FullName, rj.Username, rj.Email, rj.Password, "user")

			// Buscar o user dnv
			err = db.DB.Where("username = ? or email = ?", rj.Username, rj.Email).First(&user).Error

			// Gerar token e redirecionar o usuario
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"user_id":    user.ID,
				"expiration": time.Now().Add(time.Hour * 168).Local().Unix(), // 7 dias
				"role":       user.Role,
				"remember":   false,
			})

			tokenString, _ := token.SignedString(utils.JwtSecrect)

			// setar os cookies
			ctx.SetCookie("token", tokenString, 604800, "/", "", false, true) // cookies validos ate 7 dias

			// enviar pro cliente
			ctx.JSON(http.StatusOK, gin.H{
				"status": true,
				// "token":       tokenString,
				"redirect_to": "/dashboard",
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"error":  "Username ou email já existente",
			})
		}
	})

}

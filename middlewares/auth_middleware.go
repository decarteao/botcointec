package middlewares

import (
	"botcointec/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(role_permitted string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token_client, err := ctx.Cookie("token")

		if err != nil {
			// ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			// 	"status": false,
			// 	"error":  "Sessão inválida",
			// })
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		token, err := jwt.Parse(token_client, func(t *jwt.Token) (interface{}, error) {
			return utils.JwtSecrect, nil
		})

		if err != nil || !token.Valid {
			// ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			// 	"status": false,
			// 	"error":  "Sessão inválida",
			// })
			// fazer um flash aqui
			ctx.SetCookie("flash", "Sessão inválida", 1, "/", "", false, true)

			ctx.Redirect(http.StatusFound, "/login")
		} else {
			// verificar token
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && time.Now().Local().Unix() > int64(claims["expiration"].(float64)) && !claims["remember"].(bool) {
				// avisar do token expirado
				// ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				// 	"status": false,
				// 	"error":  "Token expirado",
				// })

				ctx.SetCookie("flash", "Sessão expirada", 1, "/", "", false, true)
				ctx.Redirect(http.StatusFound, "/login")
			} else if claims["role"] == "admin" && role_permitted == "user" {
				// redirecionar o cliente para onde deve ir
				ctx.Redirect(http.StatusFound, "/admin")
			} else if claims["role"] == "user" && role_permitted == "admin" {
				// redirecionar o cliente para onde deve ir
				ctx.Redirect(http.StatusFound, "/dashboard")
			} else {
				// token valido e login efectuado com sucesso
				ctx.Next()
			}
		}
	}
}

func LoginAndSingupForm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token_client, err := ctx.Cookie("token")

		if err != nil {
			ctx.Next()
			return
		}

		token, err := jwt.Parse(token_client, func(t *jwt.Token) (interface{}, error) {
			return utils.JwtSecrect, nil
		})

		if err != nil || !token.Valid {
			ctx.Next()
		} else {
			// verificar token
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && time.Now().Local().Unix() > int64(claims["expiration"].(float64)) && !claims["remember"].(bool) {
				ctx.Next()
			} else {
				// token valido e login efectuado com sucesso
				if claims["role"] == "admin" {
					// redirecionar o cliente para onde deve ir
					ctx.Redirect(http.StatusFound, "/admin")
				} else if claims["role"] == "user" {
					// redirecionar o cliente para onde deve ir
					ctx.Redirect(http.StatusFound, "/dashboard")
				}
			}
		}
	}
}

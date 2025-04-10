package main

import (
	"botcointec/config"
	"botcointec/controllers"
	"botcointec/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Iniciar conexao ao Banco de Dados
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	controllers.AddUser(db, "William Silva", "william", "renedcteixeira@gmail.com", "H3li0.AO26", "admin")
	controllers.AddUser(db, "Rene Decarte", "decarteao", "decartecussenha@gmail.com", "H3li0.AO26", "user")
	controllers.AddSEODefault(db)

	// Cria uma instância do Gin com middleware padrão (logger, recovery, etc)
	r := gin.Default()

	// Carrega todos os templates da pasta views e subpastas
	r.LoadHTMLGlob("views/**/*.html")
	r.Static("/static", "./static")

	// Registrar rotas index
	routes.RegistarLPRotas(r)

	// Registrar rotas admin
	routes.RegistarAdminRotas(r)

	// Registrar rotas user
	routes.RegistarUserRotas(r)

	// Registrar rotas da API
	routes.RegisterApi(r, db)

	// Roda o servidor na porta 8080
	log.Println("Website is running...")
	r.Run(":80")
}

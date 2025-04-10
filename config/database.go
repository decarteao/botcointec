package config

import (
	"botcointec/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
** Ao iniciar o sistema, deve adicionar "admin" se nao existir
** Com a senha: "H3li0.AO$$$Bot"
 */

const (
	DB_NAME = "botcointec"
	DB_USER = "root"
	DB_PASS = "qwerty1234"
)

type DataBase struct {
	DB *gorm.DB
}

func Connect() (*DataBase, error) {
	_db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_NAME)), &gorm.Config{})
	if err != nil {
		log.Println("Erro ao se conectar ao Banco de Dados:", err)
		return nil, err
	} else {
		// criar as tabelas se nao existe
		_db.AutoMigrate(&models.User{})
		_db.AutoMigrate(&models.SEO{})

		return &DataBase{DB: _db}, nil
	}
}

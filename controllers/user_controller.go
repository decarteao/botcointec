package controllers

import (
	"botcointec/config"
	"botcointec/models"
	"botcointec/utils"
	"log"
)

func AddUser(db *config.DataBase, nome string, username string, email string, password string, role string) (bool, error) {
	// adicionar na tabela: users
	// fazer md5 na password ou outra criptografia segura
	hash, err := utils.Password2Hash(password)
	if err != nil {
		return false, err
	}

	user := models.User{
		Nome:     nome,
		Email:    email,
		Username: username,
		Password: hash,
		Role:     role,
		Status:   role == "admin",
		Saldo:    0,
	}
	r := db.DB.Create(&user)
	if r.Error != nil {
		return false, r.Error
	}

	// enviar codigo de confirmacao em outro arquivo
	log.Println("Usuario criado com sucesso")
	return true, nil
}

func ConfirmarUser(email string, code string) {
	// confirmar email se tiver certo
	// se nao tiver certo, retorna false
	// se estiver certo, muda o status de pendente para confirmado e deletar o registro do codigo de confirmacao
}

func ReenviarCodigoUser(email string) {
	// so reenvie se a conta estiver no registro do codigo de confirmacao,
	// deleta o que existe e adiciona outro registro, ou apenas de um  update
}

func BuscarUser(id uint64, email string, password string) {
	// fazer md5 na password
}

func UpdatePassword(id uint64) {
	// gerar um hash aleatorio actualizar no db e enviar pro email a nova senha
}

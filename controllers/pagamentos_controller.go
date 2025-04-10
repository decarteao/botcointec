package controllers

import (
	"botcointec/config"
)

func GerarFraseMnemonica() {
	// Gerar uma nova chave mnemônica
	// entropy, err := hdwallet.GenerateEntropy()
	// if err != nil {
	// 	log.Fatalf("erro ao gerar mnemônica: %v", err)
	// }
	// log.Println("entropy:", entropy)
}
func GerarCarteira() {
	//
}
func EsperarDeposito(db *config.DataBase, montante_esperado int64, priv_key string) {
	// gerar a pub.key pelo priv.key
	// ficar dentro de um for com sleep de 5 segundos: 180 de 5 em 5
	// se confirmar o valor, muda o status no db e adiciona o saldo no user
	// -> mandar pra carteira principal
	// se expirar, muda o status no db
}

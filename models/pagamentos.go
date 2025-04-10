package models

import (
	"gorm.io/gorm"
)

type Pagamentos struct {
	gorm.Model
	tx_id    string  `json:"tx_id" gorm:"unique"`
	usdt     float64 `json:"usdt"`
	data     int64   `json:"data"`   // timestamp
	status   string  `json:"status"` // pendente | confirmado | expirado
	user_id  int64   `json:"user_id"`
	priv_key string  `json:"priv_key"` // priv-key da carteira gerada
}

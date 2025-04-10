package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nome     string  `json:"nome"`
	Email    string  `json:"email" gorm:"unique"`
	Username string  `json:"username" gorm:"unique"`
	Password string  `json:"password"`
	Role     string  `json:"role"`   // role: admin or user
	Status   bool    `json:"status"` // status: true se verificado, false se ainda
	Saldo    float64 `json:"saldo"`
}

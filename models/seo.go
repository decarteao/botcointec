package models

import "gorm.io/gorm"

type SEO struct {
	gorm.Model
	page_id        string `json:"page_id" gorm:"unique"`
	description    string `json:"description"`
	keywords       string `json:"keywords"`
	og_title       string `json:"og_title"`
	og_description string `json:"og_description"`
}

package controllers

import (
	"botcointec/config"
	"botcointec/models"
	"log"
)

func AddSEODefault(db *config.DataBase) (bool, error) {
	// adicionar seo da landing page
	seo_lp := models.SEO{}

	r := db.DB.Create(&seo_lp)
	if r.Error != nil {
		return false, r.Error
	}

	log.Println("SEO da Landing Page foi criada com sucesso")

	// adicionar seo das page de login e signup
	seo_forms := models.SEO{}

	r = db.DB.Create(&seo_forms)
	if r.Error != nil {
		return false, r.Error
	}

	log.Println("SEO dos forms foi criada com sucesso")

	return true, nil
}

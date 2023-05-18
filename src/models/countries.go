package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name       string `json:"name"`
	Flag       string `json:"flag"`
	Population int    `json:"population"`
	Coin       string `json:"coin"`
}

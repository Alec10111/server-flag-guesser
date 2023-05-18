package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name       string `json:"name"`
	Code       string `json:"code"`
	Flag       string `json:"flag"`
	Population int    `json:"population"`
	Coin       string `json:"coin"`
}

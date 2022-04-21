package models

type country struct {
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Description string `json:"description"`
	Population  int    `json:"population"`
	Coin        string `json:"coin"`
}

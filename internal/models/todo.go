package models

type Todo struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	IsChecked bool   `json:"isChecked"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

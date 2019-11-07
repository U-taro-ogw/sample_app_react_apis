package models

type Todo struct {
	Id uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status uint64 `json:"status,string"`
}

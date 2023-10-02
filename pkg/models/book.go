package models

type Book struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	NrPages int    `json:"nrPages"`
	Desc    string `json:"desc"`
}

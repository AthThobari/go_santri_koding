package models

type Post struct {
Id int `json:"id" gorm:"primary_key"`
Title string `json:"title"`
Content string `json:"content"`
}

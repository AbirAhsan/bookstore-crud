package models

import "time"

type BookModel struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
	Content       string    `json:"content" gorm:"type:text"`
}

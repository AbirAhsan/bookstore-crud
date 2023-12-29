package models

type BookModel struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"publishedYear"`
	Content       string `json:"content" gorm:"type:text"`
}

type CreateBookModel struct {
	Title         string `form:"title" binding:"required"`
	Author        string `form:"author" binding:"required"`
	PublishedYear int    `form:"publishedYear"`
	Content       string `form:"content" binding:"required"`
}

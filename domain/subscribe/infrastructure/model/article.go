package model

import "time"

type Article struct {
	ID          int    `gorm:"primary_key;column:id"`
	Title       string `gorm:"column:title"`
	Content     string `gorm:"column:content"`
	OriginalUrl string `gorm:"column:url"`
}

func NewArticle(
	id int,
	sourceId int,
	title,
	caption,
	content,
	originalUrl string,
	postTime time.Time,
) Article {
	return Article{
		ID:          id,
		Title:       title,
		Content:     content,
		OriginalUrl: originalUrl,
	}
}

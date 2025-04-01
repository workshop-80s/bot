package model

type Article struct {
	ID          int    `gorm:"primary_key;column:id"`
	Url         string `gorm:"column:url"`
	Title       string `gorm:"column:title"`
	Sapo        string `gorm:"column:sapo"`
	OriginalUrl string `gorm:"column:original_url"`
}

func (Article) TableName() string {
	return "post"
}

func NewArticle(
	id int,
	url,
	title,
	sapo,
	originalUrl string,
) Article {
	return Article{
		ID:          id,
		Url:         url,
		Title:       title,
		Sapo:        sapo,
		OriginalUrl: originalUrl,
	}
}

package model

type Article struct {
	ID    int    `gorm:"primary_key;column:id"`
	Title string `gorm:"column:title"`
	Sapo  string `gorm:"column:sapo"`
}

func (Article) TableName() string {
	return "article"
}

func NewArticle(
	id int,
	title,
	sapo string,
) Article {
	return Article{
		ID:    id,
		Title: title,
		Sapo:  sapo,
	}
}

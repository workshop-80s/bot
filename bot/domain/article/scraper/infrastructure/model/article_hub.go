package model

type ArticleHub struct {
	ID     int    `gorm:"primary_key;column:id"`
	Mode   int    `gorm:"column:mode"`
	Code   string `gorm:"column:code"`
	Domain string `gorm:"column:domain"`
}

func (ArticleHub) TableName() string {
	return "article_hub"
}

func NewArticleHub(
	id int,
	code string,
	domain string,
) ArticleHub {
	return ArticleHub{
		ID:     id,
		Code:   code,
		Domain: domain,
	}
}

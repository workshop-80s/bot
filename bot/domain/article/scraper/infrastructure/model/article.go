package model

type Article struct {
	ID           int    `gorm:"primary_key;column:id"`
	Mode         int    `gorm:"column:mode"`
	Title        string `gorm:"column:title"`
	Sapo         string `gorm:"column:sapo"`
	Content      string `gorm:"column:content"`
	Image        string `gorm:"column:image"`
	Origin       string `gorm:"column:origin"`
	ArticleHubID int    `gorm:"column:article_hub_id"`
}

func (Article) TableName() string {
	return "article"
}

func NewArticle(
	id int,
	mode int,
	title string,
	sapo string,
	content string,
	image string,
	origin string,
	articleHubId int,
) Article {
	return Article{
		ID:           id,
		Mode:         mode,
		Title:        title,
		Sapo:         sapo,
		Content:      content,
		Image:        image,
		Origin:       origin,
		ArticleHubID: articleHubId,
	}
}

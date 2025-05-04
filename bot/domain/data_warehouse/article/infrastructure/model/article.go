package model

type Article struct {
	ID           int    `gorm:"primary_key;column:id"`
	Mode         int    `gorm:"column:mode"`
	Title        string `gorm:"column:title"`
	Sapo         string `gorm:"column:sapo"`
	Content      string `gorm:"column:content"`
	Image        string `gorm:"column:image"`
	PublishedAt  string `gorm:"column:published_at"`
	Origin       string `gorm:"column:origin"`
	ArticleHubID int    `gorm:"column:article_hub_id"`
}

func (Article) TableName() string {
	return "dw_article"
}

func NewArticle(
	id int,
	mode int,
	title string,
	sapo string,
	content string,
	image string,
	publishedAt string,
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
		PublishedAt:  publishedAt,
		Origin:       origin,
		ArticleHubID: articleHubId,
	}
}

package entity

type Article struct {
	id           int
	mode         int
	title        string
	sapo         string
	content      string
	image        string
	origin       string
	articleHubId int
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
		id:           id,
		mode:         mode,
		title:        title,
		sapo:         sapo,
		content:      content,
		image:        image,
		origin:       origin,
		articleHubId: articleHubId,
	}
}

func (a Article) ID() int {
	return a.id
}

func (a Article) Mode() int {
	return a.mode
}

func (a Article) Title() string {
	return a.title
}

func (a Article) Sapo() string {
	return a.sapo
}

func (a Article) Content() string {
	return a.content
}

func (a Article) Image() string {
	return a.image
}

func (a Article) Url() string {
	return a.origin
}

func (a Article) ArticleHubID() int {
	return a.articleHubId
}

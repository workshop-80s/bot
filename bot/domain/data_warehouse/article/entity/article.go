package entity

type Article struct {
	id          int
	mode        int
	title       string
	sapo        string
	content     string
	image       string
	publishedAt string
}

func NewArticle(
	id int,
	mode int,
	title string,
	sapo string,
	content string,
	image string,
	publishedAt string,
) Article {
	return Article{
		id:          id,
		mode:        mode,
		title:       title,
		sapo:        sapo,
		content:     content,
		image:       image,
		publishedAt: publishedAt,
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

func (a Article) PublishedAt() string {
	return a.publishedAt
}

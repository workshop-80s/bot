package entity

type Article struct {
	id    int
	title string
	sapo  string
	url   string
}

func NewArticle(
	id int,
	title,
	sapo,
	url string,
) Article {
	return Article{
		id:    id,
		title: title,
		sapo:  sapo,
		url:   url,
	}
}

func (a Article) ID() int {
	return a.id
}

func (a Article) Title() string {
	return a.title
}

func (a Article) Sapo() string {
	return a.sapo
}

func (a Article) Url() string {
	return a.url
}

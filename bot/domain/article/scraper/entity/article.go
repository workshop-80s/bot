package entity

type Article struct {
	id              int
	mode            int
	title           string
	sapo            string
	content         string
	image           string
	origin          string
	articleSourceId int
}

func NewArticle(
	id int,
	mode int,
	title string,
	sapo string,
	content string,
	image string,
	origin string,
	articleSourceId int,
) Article {
	return Article{
		id:              id,
		mode:            mode,
		title:           title,
		sapo:            sapo,
		content:         content,
		image:           image,
		origin:          origin,
		articleSourceId: articleSourceId,
	}
}

func (a Article) Title() string {
	return a.title
}

func (a Article) Url() string {
	return a.origin
}

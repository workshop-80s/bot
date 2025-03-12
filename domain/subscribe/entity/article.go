package entity

type Article struct {
	title string
	url string
}

func NewArticle(
	title,
	url string,
) Article {
	return Article {
		title: title,
		url: url,
	}
}

func (a Article) Url() string {
	return a.url;
}

func (a Article) Title() string {
	return a.title;
}

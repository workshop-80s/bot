package entity

type ArticleHub struct {
	id     int
	code   string
	domain string
}

func NewArticleHub(
	id int,
	code string,
	domain string,
) ArticleHub {
	return ArticleHub{
		id:     id,
		code:   code,
		domain: domain,
	}
}

func (a ArticleHub) Id() int {
	return a.id
}

func (a ArticleHub) Code() string {
	return a.code
}

func (a ArticleHub) Domain() string {
	return a.domain
}

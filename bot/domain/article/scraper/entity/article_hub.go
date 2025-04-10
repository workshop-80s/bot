package entity

type ArticleHub struct {
	id   int
	code string
}

func NewArticleHub(
	id int,
	code string,
) ArticleHub {
	return ArticleHub{
		id:   id,
		code: code,
	}
}

func (a ArticleHub) Id() int {
	return a.id
}

func (a ArticleHub) Code() string {
	return a.code
}

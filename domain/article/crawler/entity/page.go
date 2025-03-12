package entity

type Page struct {
	url string
	title string
	content string
	thumbnail string
}

func NewPage(
	url,
	title,
	content,
	thumbnail string,
) Page {
	return Page {
		url: url,
		title: title,
		content: content,
		thumbnail: thumbnail,
	}
}


func (p Page) Url() string {
	return p.url;
}

func (p Page) Title() string {
	return p.title;
}

func (p Page) Thumbnail() string {
	return p.thumbnail;
}
package entity

type Link struct {
	id    int
	hubId int
	url   string
}

func NewLink(
	id int,
	hubId int,
	url string,
) Link {
	return Link{
		id:    id,
		hubId: hubId,
		url:   url,
	}
}

func (a Link) ID() int {
	return a.id
}

func (a Link) HubID() int {
	return a.hubId
}

func (a Link) Url() string {
	return a.url
}

package entity

type Hub struct {
	id     int
	domain string
}

func NewHub(
	id int,
	domain string,
) Hub {
	return Hub{
		id:     id,
		domain: domain,
	}
}

func (a Hub) ID() int {
	return a.id
}

func (a Hub) Domain() string {
	return a.domain
}

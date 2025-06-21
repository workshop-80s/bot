package model

type Hub struct {
	ID     int    `gorm:"primary_key;column:id"`
	Mode   int    `gorm:"column:mode"`
	Domain string `gorm:"column:domain"`
}

func (Hub) TableName() string {
	return "dw_hub"
}

func NewHub(
	id int,
	domain string,
) Hub {
	return Hub{
		ID:     id,
		Domain: domain,
	}
}

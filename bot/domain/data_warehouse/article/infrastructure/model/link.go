package model

import "time"

type Link struct {
	ID        int        `gorm:"primary_key;column:id"`
	Mode      int        `gorm:"column:mode"`
	HubID     int        `gorm:"column:hub_id"`
	Url       string     `gorm:"column:url"`
	ArticleID int        `gorm:"column:article_id"`
	CrawledAt *time.Time `gorm:"column:crawled_at"`
}

func (Link) TableName() string {
	return "dw_link"
}

func NewLink(
	id int,
	mode int,
	hubId int,
	url string,
	articleID int,
	crawledAt *time.Time,
) Link {
	return Link{
		ID:        id,
		Mode:      mode,
		HubID:     hubId,
		Url:       url,
		ArticleID: articleID,
		CrawledAt: crawledAt,
	}
}

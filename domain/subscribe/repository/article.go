package repository

import (
	"github.com/jinzhu/gorm"
)

type (
	Article interface {
		ApplyTransaction(*gorm.DB)
		Find() error
	}
)

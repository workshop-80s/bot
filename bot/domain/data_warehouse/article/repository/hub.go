package repository

import (
	"bot/domain/data_warehouse/article/entity"
)

type (
	Hub interface {
		Find() []entity.Hub
	}
)

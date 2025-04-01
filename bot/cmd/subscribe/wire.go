//go:build wireinject
// +build wireinject

package subscribe

import (
	usecase "bot/usecase/subscribe"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitialFlush(storage *gorm.DB) usecase.Flush {
	panic(wire.Build(
		usecase.ProviderFlush,
	))
}

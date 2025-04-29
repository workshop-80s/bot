package handler

import (
	"bot/infrastructure/storage/database"
	"bot/lib"
)

func slack() {
	config := lib.GetEnvConfigMap("db")
	db := database.Connect(config)
	defer func() {
		database.Disconnect(db)
	}()

	env := lib.GetEnv()
	token := lib.GetConfigString("slack", env+".bot_token")
	channels := lib.GetConfigMap("slack", env+".channel")

	uc := InitialSlack(db, token, channels)
	uc.Send()
}

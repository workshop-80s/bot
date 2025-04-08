package mail

import (
	"flag"
	"os"

	"bot/infrastructure/storage/database"
	"bot/lib"
	"bot/pkg/notify/infrastructure/repository"
	"bot/pkg/notify/usecase"
)

func Register() {
	command := flag.NewFlagSet("cmd", flag.ExitOnError)

	src := command.String("src", "", "article source")

	fooEnable := command.Bool("enable", false, "enable")
	fooName := command.String("name", "", "name")

	command.Parse(os.Args[2:])
	// fmt.Println("  source:", *src)
	// fmt.Println("  enable:", *fooEnable)
	// fmt.Println("  name:", *fooName)
	// fmt.Println("  tail:", command.Args())
	execute(*src)

}

func execute(src string) {
	config := lib.GetEnvConfigMap("db")

	db := database.Connect(config)
	defer func() {
		database.Disconnect(db)
	}()

	r := repository.NewArticle(db)
	uc := usecase.NewSlack(r)
	uc.Send()
}

package subscribe

import (
	"fmt"

	"github.com/spf13/cobra"

	"bot/domain/subscribe/infrastructure/repository"
	"bot/infrastructure/storage/database"
	"bot/usecase/subscribe"
)

func AddCommand(root *cobra.Command) {
	var cmd = &cobra.Command{
		Use:   "subscribe",
		Short: "A brief description of your command",
		Long:  `subscribe article`,
		Run: func(cmd *cobra.Command, args []string) {

			// BEGIN DI
			// Connect DB

			db := database.Connect()

			r := repository.NewArticle(db)
			uc := subscribe.NewSubscribeFlush(r)
			// END DI

			defer func() {
				fmt.Println("defer")
				database.Disconnect(db)
			}()

			uc.Flush()
		},
	}

	root.AddCommand(cmd)
}

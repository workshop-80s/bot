package subscribe

import (
	"bot/infrastructure/storage/database"
	"bot/lib"

	"github.com/spf13/cobra"
)

func AddCommand(root *cobra.Command) {
	var cmd = &cobra.Command{
		Use:   "subscribe",
		Short: "A brief description of your command",
		Long:  `subscribe article`,
		Run: func(cmd *cobra.Command, args []string) {
			config := lib.GetEnvConfigMap("db")

			db := database.Connect(config)
			defer func() {
				database.Disconnect(db)
			}()

			uc := InitialFlush(db)
			uc.Flush()
		},
	}

	root.AddCommand(cmd)
}

package article

import (
	"bot/usecase/article"

	"github.com/spf13/cobra"
)

func AddCommand(root *cobra.Command) {
	var crawlCmd = &cobra.Command{
		Use:   "article",
		Short: "A brief description of your command",
		Long:  `crawl article`,
		Run: func(cmd *cobra.Command, args []string) {
			article.Crawl();
		},
	}

	root.AddCommand(crawlCmd);
}
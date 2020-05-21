package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

const author = `Md Kawser Munshi <mkawserm@gmail.com>
`

func RegisterAuthorCMD(_ *App) *cobra.Command {
	return &cobra.Command{
		Use:   "author",
		Short: "Print Flamed authors",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf(author)
		},
	}
}

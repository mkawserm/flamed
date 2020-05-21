package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

const author = `Md Kawser Munshi <mkawserm@gmail.com>
`

var authorCMD = &cobra.Command{
	Use:   "author",
	Short: "Print Flamed authors",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf(author)
	},
}

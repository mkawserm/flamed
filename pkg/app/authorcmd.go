package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

const author = `Md Kawser Munshi <mkawserm@gmail.com>
`

var AuthorCMD = &cobra.Command{
	Use:   "author",
	Short: "Print mFlamed authors",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf(author)
	},
}

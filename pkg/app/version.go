package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "v0.1.0"

var VersionCMD = &cobra.Command{
	Use:   "version",
	Short: "Print Flamed version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(version)
	},
}

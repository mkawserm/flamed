package app

import "github.com/spf13/cobra"

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run Flamed server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

package app

import (
	"github.com/spf13/cobra"
)

func RegisterRunCMD(_ *App) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run Flamed server",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

package app

import "github.com/spf13/cobra"

func InitConfigFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&configFile, "config", "", "Config file")
}

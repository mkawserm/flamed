package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filePath string

var ConfigCMD = &cobra.Command{
	Use:   "config",
	Short: "Manage config file",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println(cmd.UsageString())
	},
}

var configCreateCMD = &cobra.Command{
	Use:   "create",
	Short: "Create config file",
	Run: func(cmd *cobra.Command, _ []string) {
		if len(filePath) == 0 {
			fmt.Println("file-path can not be empty")
			return
		}

		err := viper.WriteConfigAs(filePath)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	configCreateCMD.
		Flags().
		StringVar(
			&filePath,
			"file-path",
			"",
			"Configuration file path")
	ConfigCMD.AddCommand(configCreateCMD)
}

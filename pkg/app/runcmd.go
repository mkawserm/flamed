package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RunCMD = &cobra.Command{
	Use:   "run",
	Short: "Run command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
} // Command

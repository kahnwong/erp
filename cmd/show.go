package cmd

import (
	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show items",
	Run: func(cmd *cobra.Command, args []string) {
		core.Show()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

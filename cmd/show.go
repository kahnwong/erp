package cmd

import (
	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show items",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Show()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

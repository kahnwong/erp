package cmd

import (
	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Edit()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

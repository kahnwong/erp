package cmd

import (
	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit data",
	Run: func(cmd *cobra.Command, args []string) {
		core.Edit()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

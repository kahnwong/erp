package cmd

import (
	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add item",
	Run: func(cmd *cobra.Command, args []string) {
		core.Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

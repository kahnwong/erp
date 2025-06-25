package cmd

import (
	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

func CategoryGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocomplete []string

	if len(args) == 0 {
		autocomplete = core.AppConfig.Categories
	}

	return autocomplete, cobra.ShellCompDirectiveNoFileComp
}

var addCmd = &cobra.Command{
	Use:               "add",
	Short:             "Add item",
	ValidArgsFunction: CategoryGet,
	Run: func(cmd *cobra.Command, args []string) {
		core.Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

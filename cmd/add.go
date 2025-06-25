package cmd

import (
	"fmt"

	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add item",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		fmt.Println(core.AppConfig)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

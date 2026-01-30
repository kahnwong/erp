package cmd

import (
	"os"

	"github.com/kahnwong/erp/core"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "erp",
	Short:        "For managing perishables and consumables",
	SilenceUsage: true,
	RunE:         func(cmd *cobra.Command, args []string) error { return core.Show() },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

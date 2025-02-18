package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var licensesCmd = &cobra.Command{
	Use:     "licenses",
	Short:   "Check licenses in consumed packages",
	GroupID: "review",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licensesCmd).Standalone()

	rootCmd.AddCommand(licensesCmd)
}

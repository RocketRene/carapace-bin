package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var subcommandCmd = &cobra.Command{
	Use:   "subcommand",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subcommandCmd).Standalone()

	rootCmd.AddCommand(subcommandCmd)

	carapace.Gen(subcommandCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}

package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_helpCmd).Standalone()

	branchCmd.AddCommand(branch_helpCmd)

	carapace.Gen(branch_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(branchCmd),
	)
}

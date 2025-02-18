package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var rdpCmd = &cobra.Command{
	Use:   "rdp",
	Short: "connects to machine via RDP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdpCmd).Standalone()

	rootCmd.AddCommand(rdpCmd)

	carapace.Gen(rdpCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}

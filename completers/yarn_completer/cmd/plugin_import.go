package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Download a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_importCmd).Standalone()

	pluginCmd.AddCommand(plugin_importCmd)

	carapace.Gen(pluginCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			// TODO plugin completion
			carapace.ActionValues(),
		).ToA(),
	)
}

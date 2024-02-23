package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "manually",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().String("string", "", "some string")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"string": carapace.ActionValues("one", "two", "three"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("four", "five", "six"),
	)
}

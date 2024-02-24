package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffFilesCmd = &cobra.Command{
	Use:   "diff-files",
	Short: "Compares files in the working tree and the index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffFilesCmd).Standalone()

	diffFilesCmd.Flags().BoolS("q", "q", false, "Remain silent even for nonexistent files")
	diffFilesCmd.Flags().StringS("G", "G", "", "Look for differences whose patch text contains added/removed lines that match <regex>")
	diffFilesCmd.Flags().BoolP("cc", "c", false, "This compares stage 2 (our branch), stage 3 (their branch), and the working tree file")
	diffFilesCmd.Flags().Bool("default-prefix", false, "Use the default source and destination prefixes")
	diffFilesCmd.Flags().StringP("ignore-matching-lines", "I", "", "Ignore changes whose all lines match <regex>")
	diffFilesCmd.Flags().Bool("no-relative", false, "Do not show relative pathnames")
	diffFilesCmd.Flags().Bool("no-textconv", false, "Disallow external text conversion filters to be run when comparing binary files")
	diffFilesCmd.Flags().String("relative", "", "Show relative pathnames")
	diffFilesCmd.Flags().String("rotate-to", "", "Move the files before the named <file> to the end")
	diffFilesCmd.Flags().String("skip-to", "", "Discard the files before the named <file> from the output")
	addDiffFlags(diffFilesCmd)
	addPrettyFlags(diffFilesCmd)
	rootCmd.AddCommand(diffFilesCmd)

	carapace.Gen(diffFilesCmd).FlagCompletion(carapace.ActionMap{
		"relative":  carapace.ActionValues(), // TODO
		"rotate-to": carapace.ActionValues(), // TODO
		"skip-to":   carapace.ActionValues(), // TODO
	})

	carapace.Gen(diffFilesCmd).PositionalAnyCompletion(
		git.ActionChanges(git.ChangeOpts{Unstaged: true}).FilterArgs(),
	)
}

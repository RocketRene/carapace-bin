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

	diffFilesCmd.Flags().BoolS("R", "R", false, "Swap two inputs")
	diffFilesCmd.Flags().BoolS("q", "q", false, "Remain silent even for nonexistent files")
	diffFilesCmd.Flags().BoolS("u", "u", false, "Generate patch")
	diffFilesCmd.Flags().BoolS("0", "0", false, "Omit diff output for unmerged entries")
	diffFilesCmd.Flags().StringS("O", "O", "", "Control the order in which files appear in the output")
	diffFilesCmd.Flags().StringS("G", "G", "", "Look for differences whose patch text contains added/removed lines that match <regex>")
	diffFilesCmd.Flags().StringS("S", "S", "", "Look for differences that change the number of occurrences of the specified string")
	diffFilesCmd.Flags().StringS("l", "l", "", "Prevent the exhaustive portion of rename/copy detection from running")
	diffFilesCmd.Flags().BoolS("z", "z", false, "Do not munge pathnames and use NULs as output field terminators")
	diffFilesCmd.Flags().String("abbrev", "", "Show the shortest prefix that is at least <n> hexdigits long")
	diffFilesCmd.Flags().String("anchored", "", "Generate a diff using the \"anchored diff\" algorithm")
	diffFilesCmd.Flags().BoolP("base", "1", false, "Diff against the \"base\" version")
	diffFilesCmd.Flags().Bool("binary", false, "Output a binary diff that can be applied with git-apply")
	diffFilesCmd.Flags().StringP("break-rewrites", "B", "", "Break complete rewrite changes into pairs of delete and create")
	diffFilesCmd.Flags().BoolP("cc", "c", false, "This compares stage 2 (our branch), stage 3 (their branch), and the working tree file")
	diffFilesCmd.Flags().Bool("check", false, "Warn if changes introduce conflict markers or whitespace errors")
	diffFilesCmd.Flags().String("color", "", "Show colored diff")
	diffFilesCmd.Flags().String("color-moved", "", "Moved lines of code are colored differently")
	diffFilesCmd.Flags().String("color-moved-ws", "", "This configures how whitespace is ignored when performing the move detection")
	diffFilesCmd.Flags().String("color-words", "", "Equivalent to --word-diff=color plus --word-diff-regex=<regex>")
	diffFilesCmd.Flags().Bool("compact-summary", false, "Output a condensed summary of extended header information")
	diffFilesCmd.Flags().Bool("cumulative", false, "Synonym for --dirstat=cumulative")
	diffFilesCmd.Flags().Bool("default-prefix", false, "Use the default source and destination prefixes")
	diffFilesCmd.Flags().String("diff-algorithm", "", "Choose a diff algorithm")
	diffFilesCmd.Flags().String("diff-filter", "", "Select only files matching given filter")
	diffFilesCmd.Flags().StringP("dirstat", "X", "", "Output the distribution of relative amount of changes for each sub-directory")
	diffFilesCmd.Flags().String("dirstat-by-file", "", "Synonym for --dirstat=files,param1,param2...")
	diffFilesCmd.Flags().String("dst-prefix", "", "Show the given destination prefix instead of \"b/\"")
	diffFilesCmd.Flags().Bool("exit-code", false, "Make the program exit with codes similar to diff")
	diffFilesCmd.Flags().Bool("ext-diff", false, "Allow an external diff helper to be executed")
	diffFilesCmd.Flags().StringP("find-copies", "C", "", "Detect copies as well as renames")
	diffFilesCmd.Flags().Bool("find-copies-harder", false, "Inspect unmodified files as candidates for the source of copy")
	diffFilesCmd.Flags().String("find-object", "", "Look for differences that change the number of occurrences of the specified object")
	diffFilesCmd.Flags().StringP("find-renames", "M", "", "Detect renames")
	diffFilesCmd.Flags().Bool("full-index", false, "Show the full pre- and post-image blob object names")
	diffFilesCmd.Flags().BoolP("function-context", "W", false, "Show whole function as context lines for each change")
	diffFilesCmd.Flags().Bool("histogram", false, "Generate a diff using the \"histogram diff\" algorithm")
	diffFilesCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines")
	diffFilesCmd.Flags().Bool("ignore-blank-lines", false, "Ignore changes whose lines are all blank")
	diffFilesCmd.Flags().Bool("ignore-cr-at-eol", false, "Ignore carriage-return at the end of line when doing a comparison")
	diffFilesCmd.Flags().StringP("ignore-matching-lines", "I", "", "Ignore changes whose all lines match <regex>")
	diffFilesCmd.Flags().Bool("ignore-space-at-eol", false, "Ignore changes in whitespace at EOL")
	diffFilesCmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in amount of whitespace")
	diffFilesCmd.Flags().String("ignore-submodules", "", "Ignore changes to submodules in the diff generation")
	diffFilesCmd.Flags().Bool("indent-heuristic", false, "Enable the heuristic that shifts diff hunk boundaries")
	diffFilesCmd.Flags().String("inter-hunk-context", "", "Show the context between diff hunks, up to the specified number of lines")
	diffFilesCmd.Flags().BoolP("irreversible-delete", "D", false, "Omit the preimage for deletes")
	diffFilesCmd.Flags().Bool("ita-invisible-in-index", false, "By default entries added by \"git add -N\" appear as an existing empty file")
	diffFilesCmd.Flags().String("line-prefix", "", "Prepend an additional prefix to every line of output")
	diffFilesCmd.Flags().Bool("minimal", false, "Spend extra time to make sure the smallest possible diff is produced")
	diffFilesCmd.Flags().Bool("name-only", false, "Show only names of changed files")
	diffFilesCmd.Flags().Bool("name-status", false, "Show only names and status of changed files")
	diffFilesCmd.Flags().Bool("no-color", false, "Turn off colored diff")
	diffFilesCmd.Flags().Bool("no-color-moved", false, "Turn off move detection")
	diffFilesCmd.Flags().Bool("no-color-moved-ws", false, "Do not ignore whitespace when performing move detection")
	diffFilesCmd.Flags().Bool("no-ext-diff", false, "Disallow external diff drivers")
	diffFilesCmd.Flags().String("no-indent-heuristic", "", "the indent heuristic.")
	diffFilesCmd.Flags().BoolP("no-patch", "s", false, "Suppress all output from the diff machinery")
	diffFilesCmd.Flags().Bool("no-prefix", false, "Do not show any source or destination prefix")
	diffFilesCmd.Flags().Bool("no-relative", false, "Do not show relative pathnames")
	diffFilesCmd.Flags().Bool("no-rename-empty", false, "Whether to use empty blobs as rename source")
	diffFilesCmd.Flags().Bool("no-renames", false, "Turn off rename detection")
	diffFilesCmd.Flags().Bool("no-textconv", false, "Disallow external text conversion filters to be run when comparing binary files")
	diffFilesCmd.Flags().Bool("numstat", false, "Show number of added and deleted lines in decimal notation")
	diffFilesCmd.Flags().BoolP("ours", "2", false, "Diff against \"our branch\"")
	diffFilesCmd.Flags().String("output", "", "Output to a specific file instead of stdout")
	diffFilesCmd.Flags().String("output-indicator-context", "", "Specify the character used to indicate context lines")
	diffFilesCmd.Flags().String("output-indicator-new", "", "Specify the character used to indicate new lines")
	diffFilesCmd.Flags().String("output-indicator-old", "", "Specify the character used to indicate old lines")
	diffFilesCmd.Flags().BoolP("patch", "p", false, "Generate patch")
	diffFilesCmd.Flags().Bool("patch-with-raw", false, "Synonym for -p --raw")
	diffFilesCmd.Flags().Bool("patch-with-stat", false, "Synonym for -p --stat")
	diffFilesCmd.Flags().Bool("patience", false, "Generate a diff using the \"patience diff\" algorithm")
	diffFilesCmd.Flags().Bool("pickaxe-all", false, "When -S or -G finds a change, show all the changes in that changeset")
	diffFilesCmd.Flags().Bool("pickaxe-regex", false, "Treat the <string> given to -S as an extended POSIX regular expression to match")
	diffFilesCmd.Flags().Bool("quiet", false, "Disable all output of the program")
	diffFilesCmd.Flags().Bool("raw", false, "Generate the diff in raw format")
	diffFilesCmd.Flags().String("relative", "", "Show relative pathnames")
	diffFilesCmd.Flags().Bool("rename-empty", false, "Whether to use empty blobs as rename source")
	diffFilesCmd.Flags().String("rotate-to", "", "Move the files before the named <file> to the end")
	diffFilesCmd.Flags().Bool("shortstat", false, "Output only the last line of the --stat format")
	diffFilesCmd.Flags().String("skip-to", "", "Discard the files before the named <file> from the output")
	diffFilesCmd.Flags().String("src-prefix", "", "Show the given source prefix instead of \"a/\"")
	diffFilesCmd.Flags().String("stat", "", "Generate a diffstat")
	diffFilesCmd.Flags().String("submodule", "", "Specify how differences in submodules are shown")
	diffFilesCmd.Flags().Bool("summary", false, "Output a condensed summary of extended header information")
	diffFilesCmd.Flags().BoolP("text", "a", false, "Treat all files as text")
	diffFilesCmd.Flags().Bool("textconv", false, "Allow external text conversion filters to be run when comparing binary files")
	diffFilesCmd.Flags().BoolP("theirs", "3", false, "Diff against \"their branch\"")
	diffFilesCmd.Flags().StringP("unified", "U", "", "Generate diffs with <n> lines of context instead of the usual three")
	diffFilesCmd.Flags().String("word-diff", "", "Show a word diff, using the <mode> to delimit changed words")
	diffFilesCmd.Flags().String("word-diff-regex", "", "Use <regex> to decide what a word is")
	diffFilesCmd.Flags().String("ws-error-highlight", "", "Highlight whitespace errors in the context")
	rootCmd.AddCommand(diffFilesCmd)

	diffFilesCmd.Flag("color").NoOptDefVal = " "

	carapace.Gen(diffFilesCmd).FlagCompletion(carapace.ActionMap{
		"color":                    git.ActionColorModes(),
		"color-moved":              git.ActionColorMovedModes(),
		"color-moved-ws":           git.ActionColorMovedWsModes(),
		"diff-algorithm":           git.ActionDiffAlgorithms(),
		"diff-filter":              git.ActionDiffFilters().UniqueList(","),
		"dirstat":                  carapace.ActionValues(),
		"dirstat-by-file":          carapace.ActionValues(),
		"dst-prefix":               carapace.ActionValues(),
		"find-copies":              carapace.ActionValues(),
		"find-object":              carapace.ActionValues(),
		"find-renames":             carapace.ActionValues(),
		"ignore-matching-lines":    carapace.ActionValues(),
		"ignore-submodules":        carapace.ActionValues(),
		"inter-hunk-context":       carapace.ActionValues(),
		"line-prefix":              carapace.ActionValues(),
		"no-indent-heuristic":      carapace.ActionValues(),
		"output":                   carapace.ActionValues(),
		"output-indicator-context": carapace.ActionValues(),
		"output-indicator-new":     carapace.ActionValues(),
		"output-indicator-old":     carapace.ActionValues(),
		"relative":                 carapace.ActionValues(),
		"rotate-to":                carapace.ActionValues(),
		"skip-to":                  carapace.ActionValues(),
		"src-prefix":               carapace.ActionValues(),
		"stat":                     carapace.ActionValues(),
		"submodule":                carapace.ActionValues(),
		"unified":                  carapace.ActionValues(),
		"word-diff":                carapace.ActionValues(),
		"word-diff-regex":          carapace.ActionValues(),
		"ws-error-highlight":       carapace.ActionValues(),
	})

	carapace.Gen(diffFilesCmd).PositionalAnyCompletion(
		git.ActionChanges(git.ChangeOpts{Unstaged: true}).FilterArgs(),
	)
}

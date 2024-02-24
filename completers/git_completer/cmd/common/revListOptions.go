package common

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

func AddRevListOptions(cmd *cobra.Command) {
	// TODO a lot of these flags are repeatable
	cmd.Flags().Bool("abbrev-commit", false, "Instead of showing the full 40-byte hexadecimal commit object name")
	cmd.Flags().String("after", "", "Show commits more recent than a specific date")
	cmd.Flags().Bool("all", false, "Pretend as if all the refs in refs/ are listed on the command line as <commit>")
	cmd.Flags().Bool("all-match", false, "Limit the commits output to ones that match all given --grep")
	cmd.Flags().Bool("alternate-refs", false, "Pretend as if all objects mentioned as ref tips of alternate repositories were listed on the command line")
	cmd.Flags().String("ancestry-path", "", "Limit the displayed commits to those which are an ancestor of <commit>")
	cmd.Flags().String("author", "", "Limit the commits output to ones with author")
	cmd.Flags().Bool("author-date-order", false, "Show no parents before all of its children are shown")
	cmd.Flags().Bool("basic-regexp", false, "Consider the limiting patterns to be basic regular expressions")
	cmd.Flags().String("before", "", "Show commits older than a specific date")
	cmd.Flags().Bool("bisect", false, "Limit output to the one commit object which is roughly halfway between included and excluded commits")
	cmd.Flags().Bool("bisect-all", false, "This outputs all the commit objects between the included and excluded commits")
	cmd.Flags().Bool("bisect-vars", false, "This calculates the same as --bisect, except that refs in refs/bisect/ are not used")
	cmd.Flags().Bool("boundary", false, "Output excluded boundary commits")
	cmd.Flags().String("branches", "", "Pretend as if all the refs in refs/heads are listed on the command line as <commit>")
	cmd.Flags().Bool("cherry", false, "A synonym for --right-only --cherry-mark --no-merges")
	cmd.Flags().Bool("cherry-mark", false, "Like --cherry-pick but mark equivalent commits with = rather than omitting them")
	cmd.Flags().Bool("cherry-pick", false, "Omit any commit that introduces the same change as another commit")
	cmd.Flags().Bool("children", false, "Print also the children of the commit")
	cmd.Flags().Bool("commit-header", false, "Overrides a previous --no-commit-header")
	cmd.Flags().String("committer", "", "Limit the commits output to ones with committer")
	cmd.Flags().Bool("count", false, "Print a number stating how many commits would have been listed")
	cmd.Flags().String("date", "", "Only takes effect for dates shown in human-readable format")
	cmd.Flags().Bool("date-order", false, "Show no parents before all of its children are shown")
	cmd.Flags().Bool("dense", false, "Commits that are walked are included if they are not TREESAME to any parent")
	cmd.Flags().String("disk-usage", "", "Suppress normal output")
	cmd.Flags().Bool("do-walk", false, "Overrides a previous --no-walk")
	cmd.Flags().String("encoding", "", "Commit objects record the character encoding used for the log message in their encoding header")
	cmd.Flags().String("exclude", "", "Do not include refs matching <glob-pattern>")
	cmd.Flags().String("exclude-first-parent-only", "", "Follow only the first parent commit upon seeing a merge commit")
	cmd.Flags().String("exclude-hidden", "", "Do not include refs that would be hidden by git-receive-pack or git-upload-pack")
	cmd.Flags().Bool("exclude-promisor-objects", false, "Prefilter object traversal at promisor boundary")
	cmd.Flags().String("expand-tabs", "", "Perform a tab expansion")
	cmd.Flags().BoolP("extended-regexp", "E", false, "Consider the limiting patterns to be extended regular expressions")
	cmd.Flags().String("filter", "", "Omits objects from the list of printed objects")
	cmd.Flags().Bool("filter-print-omitted", false, "Prints a list of the objects omitted by the filter")
	cmd.Flags().Bool("filter-provided-objects", false, "Filter the list of explicitly provided objects")
	cmd.Flags().Bool("first-parent", false, "Follow only the first parent commit upon seeing a merge commit")
	cmd.Flags().BoolP("fixed-strings", "F", false, "Consider the limiting patterns to be fixed strings")
	cmd.Flags().String("format", "", "Pretty-print the contents of the commit logs in a given format")
	cmd.Flags().Bool("full-history", false, "Same as the default mode, but does not prune some history")
	cmd.Flags().String("glob", "", "Pretend as if all the refs matching shell glob <glob-pattern> are listed on the command line as <commit>")
	cmd.Flags().Bool("graph", false, "Draw a text-based graphical representation of the commit history")
	cmd.Flags().String("grep", "", "Limit the commits output to ones with log message that matches the specified pattern")
	cmd.Flags().String("grep-reflog", "", "Limit the commits output to ones with reflog entries that match the specified pattern")
	cmd.Flags().Bool("header", false, "Print the contents of the commit in raw-format")
	cmd.Flags().Bool("ignore-missing", false, "Ignore invalid object names in the input")
	cmd.Flags().Bool("in-commit-order", false, "Print tree and blob ids in order of the commits")
	cmd.Flags().Bool("indexed-objects", false, "Pretend as if all trees and blobs used by the index are listed on the command line")
	cmd.Flags().Bool("invert-grep", false, "Limit the commits output to ones with log message that do not match the pattern specified with --grep")
	cmd.Flags().Bool("left-only", false, "List only commits on the left side of a symmetric difference")
	cmd.Flags().Bool("left-right", false, "Mark which side of a symmetric difference a commit is reachable from")
	cmd.Flags().String("max-age", "", "Limit the commits output to specified time range")
	cmd.Flags().StringP("max-count", "n", "", "Limit the number of commits to output")
	cmd.Flags().String("max-parents", "", "Show only commits which have at most that many parent commits")
	cmd.Flags().Bool("merge", false, "Show refs that touch files having a conflict and don’t exist on all heads to merge")
	cmd.Flags().Bool("merges", false, "Print only merge commits")
	cmd.Flags().String("min-age", "", "Limit the commits output to specified time range")
	cmd.Flags().String("min-parents", "", "Show only commits which have at least that many parent commits")
	cmd.Flags().String("missing", "", "A debug option to help with future \"partial clone\" development")
	cmd.Flags().Bool("no-abbrev-commit", false, "Show the full 40-byte hexadecimal commit object name")
	cmd.Flags().Bool("no-commit-header", false, "Suppress the header line containing \"commit\" and the object ID printed before the specified format")
	cmd.Flags().Bool("no-expand-tabs", false, "Do not perform a tab expansion")
	cmd.Flags().Bool("no-filter", false, "Turn off any previous --filter= argument")
	cmd.Flags().Bool("no-max-parents", false, "Reset maximum parents limit")
	cmd.Flags().Bool("no-merges", false, "Do not print commits with more than one parent")
	cmd.Flags().Bool("no-min-parents", false, "Reset minimum parents limit")
	cmd.Flags().Bool("no-object-names", false, "Does not print the names of the object IDs that are found")
	cmd.Flags().String("no-walk", "", "Only show the given commits, but do not traverse their ancestors")
	cmd.Flags().Bool("not", false, "Reverses the meaning of the ^ prefix (or lack thereof) for all following revision specifiers")
	cmd.Flags().Bool("object-names", false, "Print the names of the object IDs that are found")
	cmd.Flags().Bool("objects", false, "Print the object IDs of any object referenced by the listed commits")
	cmd.Flags().Bool("objects-edge", false, "Similar to --objects, but also print the IDs of excluded commits prefixed with a “-” character")
	cmd.Flags().Bool("objects-edge-aggressive", false, "Similar to --objects-edge, but it tries harder to find excluded commits")
	cmd.Flags().Bool("oneline", false, "This is a shorthand for \"--pretty=oneline --abbrev-commit\" used together")
	cmd.Flags().Bool("parents", false, "Print also the parents of the commit")
	cmd.Flags().BoolP("perl-regexp", "P", false, "Consider the limiting patterns to be Perl-compatible regular expressions")
	cmd.Flags().String("pretty", "", "Pretty-print the contents of the commit logs in a given format")
	cmd.Flags().String("progress", "", "Show progress reports on stderr as objects are considered")
	cmd.Flags().Bool("quiet", false, "Don’t print anything to standard output")
	cmd.Flags().Bool("reflog", false, "Pretend as if all objects mentioned by reflogs are listed on the command line as <commit>")
	cmd.Flags().BoolP("regexp-ignore-case", "i", false, "Match the regular expression limiting patterns without regard to letter case")
	cmd.Flags().Bool("relative-date", false, "Synonym for --date=relative")
	cmd.Flags().String("remotes", "", "Pretend as if all the refs in refs/remotes are listed on the command line as <commit>")
	cmd.Flags().Bool("remove-empty", false, "Stop when a given path disappears from the tree")
	cmd.Flags().Bool("reverse", false, "Output the commits chosen to be shown in reverse order")
	cmd.Flags().Bool("right-only", false, "List only commits on the left side of a symmetric difference")
	cmd.Flags().String("show-linear-break", "", "Put a barrier in between linear branches")
	cmd.Flags().Bool("show-pulls", false, "Include all commits from the default mode, but also any merge commits")
	cmd.Flags().Bool("show-signature", false, "Check the validity of a signed commit object by passing the signature to gpg --verify")
	cmd.Flags().Bool("simplify-by-decoration", false, "Commits that are referred by some branch or tag are selected")
	cmd.Flags().Bool("simplify-merges", false, "Simplify each commit C to its replacement C' in the final history")
	cmd.Flags().String("since", "", "Show commits more recent than a specific date")
	cmd.Flags().String("since-as-filter", "", "Show all commits more recent than a specific date")
	cmd.Flags().Bool("single-worktree", false, "Examine the current working tree only")
	cmd.Flags().String("skip", "", "Skip number commits before starting to show the commit output")
	cmd.Flags().Bool("sparse", false, "All commits in the simplified history are shown")
	cmd.Flags().Bool("stdin", false, "Read <commit> additionally from the standard input")
	cmd.Flags().String("tags", "", "Pretend as if all the refs in refs/tags are listed on the command line as <commit>")
	cmd.Flags().Bool("timestamp", false, "Print the raw commit timestamp")
	cmd.Flags().Bool("topo-order", false, "Show no parents before all of its children are shown")
	cmd.Flags().Bool("unpacked", false, "Print the object IDs that are not in packs")
	cmd.Flags().String("until", "", "Show commits older than a specific date")
	cmd.Flags().Bool("use-bitmap-index", false, "Try to speed up the traversal using the pack bitmap index")
	cmd.Flags().BoolP("walk-reflogs", "g", false, "Walk reflog entries from the most recent one to older ones")

	cmd.Flag("ancestry-path").NoOptDefVal = " "
	cmd.Flag("ancestry-path").NoOptDefVal = " "
	cmd.Flag("branches").NoOptDefVal = " "
	cmd.Flag("disk-usage").NoOptDefVal = " "
	cmd.Flag("exclude-hidden").NoOptDefVal = " "
	cmd.Flag("no-walk").NoOptDefVal = " "
	cmd.Flag("pretty").NoOptDefVal = " "
	cmd.Flag("remotes").NoOptDefVal = " "
	cmd.Flag("show-linear-break").NoOptDefVal = " "
	cmd.Flag("tags").NoOptDefVal = " "

	// TODO still a lot of completions missing / incorrect
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"after":     time.ActionDate(),
		"author":    git.ActionAuthors(),
		"before":    time.ActionDate(),
		"branches":  git.ActionLocalBranches(),
		"committer": git.ActionCommitters(),
		"date": carapace.ActionValuesDescribed(
			"relative", "shows dates relative to the current time",
			"local", "is an alias for --date=default-local",
			"iso", "shows timestamps in a ISO 8601-like format",
			"iso8601", "shows timestamps in a ISO 8601-like format",
			"iso-strict", "shows timestamps in strict ISO 8601 format",
			"iso8601-strict", "shows timestamps in strict ISO 8601 format",
			"rfc", "shows timestamps in RFC 2822 format",
			"rfc2822", "shows timestamps in RFC 2822 format",
			"short", "shows only the date, but not the time, in YYYY-MM-DD format",
			"raw", "shows the date as seconds since the epoch (1970-01-01 00:00:00 UTC)",
			"human", "shows the date in human readable format",
			"unix", "shows the date as a Unix epoch timestamp (seconds since 1970)",
			"format:", "feeds the given format to your system strftime",
			"format:%c", "show the date in your system locale’s preferred format",
			"default", "the default format",
		),
		"disk-usage": carapace.ActionValues("human"),
		"exclude": carapace.Batch(
			git.ActionLocalBranches(),
			git.ActionTags(),
			git.ActionRemotes(),
		).ToA(),
		"exclude-hidden": carapace.ActionValues("receive", "uploadpack"),
		"format":         carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"missing": carapace.ActionValues(
			"error", "requests that rev-list stop with an error if a missing object is encountered",
			"allow-any", "allow object traversal to continue if a missing object is encountered",
			"allow-promisor", "is like allow-any, but will only allow object traversal to continue for EXPECTED promisor missing objects",
			"print", "is like allow-any, but will also print a list of the missing objects",
		),
		"no-walk":         carapace.ActionValues("sorted", "unsorted"),
		"pretty":          carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:"),
		"remotes":         git.ActionRemotes(),
		"since":           time.ActionDate(),
		"since-as-filter": time.ActionDate(),
		"tags":            git.ActionTags(),
		"until":           time.ActionDate(),
	})
}

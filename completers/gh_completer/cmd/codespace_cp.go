package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace/pkg/util"
	"github.com/spf13/cobra"
)

var codespace_cpCmd = &cobra.Command{
	Use:   "cp [-e] [-r] [-- [<scp flags>...]] <sources>... <dest>",
	Short: "Copy files between local and remote file systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_cpCmd).Standalone()

	codespace_cpCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_cpCmd.Flags().BoolP("expand", "e", false, "Expand remote file names on remote shell")
	codespace_cpCmd.Flags().StringP("profile", "p", "", "Name of the SSH profile to use")
	codespace_cpCmd.Flags().BoolP("recursive", "r", false, "Recursively copy directories")
	codespace_cpCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_cpCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespaceCmd.AddCommand(codespace_cpCmd)

	// TODO profile completion
	carapace.Gen(codespace_cpCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(codespace_cpCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO remove haspathprefix
			if util.HasPathPrefix(c.Value) {
				return carapace.ActionFiles()
			} else if !strings.HasPrefix(c.Value, "remote:/") {
				return carapace.ActionValues("remote:/").NoSpace()
			}

			c.Value = strings.TrimPrefix(c.Value, "remote:")
			return action.ActionCodespacePath(
				codespace_cpCmd.Flag("codespace").Value.String(),
				codespace_cpCmd.Flag("expand").Changed,
			).Invoke(c).Prefix("remote:").ToA()
		}),
	)
}

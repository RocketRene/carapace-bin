package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_unlockCmd = &cobra.Command{
	Use:     "unlock {<number> | <url>}",
	Short:   "Unlock pull request conversation",
	GroupID: "targeted",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_unlockCmd).Standalone()

	prCmd.AddCommand(pr_unlockCmd)

	carapace.Gen(pr_unlockCmd).PositionalCompletion(
		action.ActionPullRequests(pr_lockCmd, action.PullRequestOpts{Open: true, Closed: true, Merged: true}), // TODO complete only locked pr
	)
}

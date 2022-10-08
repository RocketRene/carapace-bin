package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_clusterrolebindingCmd = &cobra.Command{
	Use:   "clusterrolebinding",
	Short: "Create a cluster role binding for a particular cluster role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_clusterrolebindingCmd).Standalone()
	create_clusterrolebindingCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_clusterrolebindingCmd.Flags().String("clusterrole", "", "ClusterRole this ClusterRoleBinding should reference")
	create_clusterrolebindingCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_clusterrolebindingCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_clusterrolebindingCmd.Flags().StringArray("group", []string{}, "Groups to bind to the clusterrole. The flag can be repeated to add multiple groups.")
	create_clusterrolebindingCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_clusterrolebindingCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_clusterrolebindingCmd.Flags().StringArray("serviceaccount", []string{}, "Service accounts to bind to the clusterrole, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.")
	create_clusterrolebindingCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_clusterrolebindingCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_clusterrolebindingCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_clusterrolebindingCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_clusterrolebindingCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_clusterrolebindingCmd)

	carapace.Gen(create_clusterrolebindingCmd).FlagCompletion(carapace.ActionMap{
		"clusterrole":    action.ActionResources("", "clusterrole"),
		"dry-run":        action.ActionDryRunModes(),
		"output":         action.ActionOutputFormats(),
		"serviceaccount": action.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
		"validate":       kubectl.ActionValidationModes(),
	})
}

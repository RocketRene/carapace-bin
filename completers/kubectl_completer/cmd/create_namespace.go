package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Short:   "Create a namespace with the specified name",
	Aliases: []string{"ns"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_namespaceCmd).Standalone()
	create_namespaceCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_namespaceCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_namespaceCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_namespaceCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_namespaceCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_namespaceCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_namespaceCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_namespaceCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_namespaceCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_namespaceCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_namespaceCmd)

	carapace.Gen(create_namespaceCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}

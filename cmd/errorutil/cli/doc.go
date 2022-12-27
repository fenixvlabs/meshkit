package cli

import "github.com/spf13/cobra"

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Print the documentation",
	Long:  "Print the documentation",
	Run: func(cmd *cobra.Command, args []string) {
		println(`
This tool analyzes, verifies and updates MeshKit compatible errors in Meshery Go source code trees.

A MeshKit compatible error consist of
- An error code defined as a constant or variable (preferably constant), of type string.
  - The naming convention for these variables is the regex "^Err[A-Z].+Code$", e.g. ErrApplyManifestCode.
  - The initial value of the code is a placeholder string, e.g. "replace_me", set by the developer.
  - The final value of the code is an integer, set by this tool, as part of a CI workflow.
  - Error details defined using the function errors.New(code, severity, sdescription, ldescription, probablecause, remedy) from MeshKit.
  - The first parameter, 'code', has to be passed as the error code constant (or variable), not a string literal.
  - The second parameter, 'severity', has its own type; consult its Go-doc for further details.
  - The remaining parameters are string arrays for short and long description, probable cause, and suggested remediation.
  - Use string literals in these string arrays, not constants or variables, for any static texts.
  - Capitalize the first letter of each statement.
  - Call expressions can be used but will be ignored by the tool when exporting error details for the documentation.
  - Do not concatenate strings using the '+' operator, just add multiple elements to the string array.

Additionally, the following conventions apply:
  - Errors are defined in each package, in a file named error.go
  - Errors are namespaced to components, i.e. they need to be unique within a component (see below).
  - Errors are not to be reused across components and modules.
  - There are no predefined error code ranges for components. Every component is free to use its own range.
  - Codes carry no meaning, as e.g. HTTP status codes do.

This tool produces three files:
  - errorutil_analyze_errors.json: raw data with all errors and some metadata
  - errorutil_analyze_summary.json: summary of raw data, also used for validation and troubleshooting
  - errorutil_errors_export.json: export of errors which can be used to create the error code reference on the Meshery website

Typically, the 'analyze' command of the tool is used by the developer to verify errors, i.e. that there are no duplicate names or details.
A CI workflow is used to replace the placeholder code strings with integer code, and export errors. Using this export, the workflow updates 
the error code reference documentation in the Meshery repository.

Meshery components and this tool:
  - Meshery components have a name and a type.
  - An example of a component is MeshKit with 'meshkit' as name, and 'library' as type.
  - Often, a specific component corresponds to one git repository.
  - The tool requires a file called component_info.json.
  This file has the following content, with concrete values specific for each component:
  {
    "name": "meshkit",
    "type": "library",
    "next_error_code": 1014
  }
  - next_error_code is the value used by the tool to replace the error code placeholder string with the next integer.
  - The tool updates next_error_code. 
`)
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"<%= importName %>/metadata"
)

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "returns <%= name %> version",
	Long:  `returns <%= name %> version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("<%= name %> v%s\n", metadata.GetVersion())
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}

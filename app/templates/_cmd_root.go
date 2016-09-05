package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Verbose determines how verbose <%= name %> will run under
var Verbose int

// RootCmd is the root command for <%= name %> CLI application
var RootCmd = &cobra.Command{
	Use:   "<%= name %>",
	Short: "<%= description %>",
	Long:  "<%= description %>",
}

// Execute runs RootCmd to initialize <%=name %> CLI application
func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	RootCmd.PersistentFlags().IntVarP(
		&Verbose, "verbose", "v", 0,
		"Verbosity level => v0: Error, v1=Warning, v2=Info, v3=Debug",
	)
}

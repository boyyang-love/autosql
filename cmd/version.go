package cmd

import (
	"autoSql-cobra/utills/getRunDir"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
	version := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v1.0.0")
			getRunDir.RunDir()
		},
		Short: "get version [autosql version]",
	}

	return version
}

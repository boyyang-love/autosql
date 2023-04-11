package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "autosql",
	Run:   rootCmdRun(),
	Short: "Auto generate SQL",
	Long:  `Auto generate SQL by Excel files`,
}

func rootCmdRun() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		//fmt.Println(args)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

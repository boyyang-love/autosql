package cmd

import (
	"autoSql-cobra/utills/setYaml"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(setSchema())
}

func setSchema() *cobra.Command {
	var s string
	schema := &cobra.Command{
		Use:   "schema",
		Short: "Set the schema for the sql",
		Run: func(cmd *cobra.Command, args []string) {
			err := setYaml.SetYaml("schema", s)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("set schema successfully")
			}
		},
	}

	schema.Flags().StringVarP(&s, "schema", "s", "", "autosql schema -s USERS")

	return schema
}

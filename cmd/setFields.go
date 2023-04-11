package cmd

import (
	"autoSql-cobra/utills/setYaml"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	RootCmd.AddCommand(setFields())
}

func setFields() *cobra.Command {
	var f string
	fields := &cobra.Command{
		Use:   "fields",
		Short: "set sql fields",
		Run: func(c *cobra.Command, args []string) {
			fields := strings.Split(f, ",")
			if len(fields) != 0 {
				err := setYaml.SetYaml("fields", fields)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("set fields successfully")
				}
			}
		},
	}

	fields.Flags().StringVarP(&f, "fields", "f", "", "sql fields -f name,age,sex")

	return fields
}

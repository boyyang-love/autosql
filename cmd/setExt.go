package cmd

import (
	"autoSql-cobra/utills/setYaml"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	RootCmd.AddCommand(setExt())
}

func setExt() *cobra.Command {
	var exts string
	ext := &cobra.Command{
		Use:   "ext",
		Short: "set files extension",
		Run: func(c *cobra.Command, args []string) {
			if len(strings.Split(exts, ".")) == 0 {
				return
			}
			err := setYaml.SetYaml("exts", strings.Split(exts, ","))
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("set extensions successfully")
			}
		},
	}

	ext.Flags().StringVarP(&exts, "ext", "e", "", "set files extension")

	return ext
}

package cmd

import (
	"autoSql-cobra/utills/resolveExcel"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(generageSql())
}

func generageSql() *cobra.Command {
	generateSql := &cobra.Command{
		Use:   "auto",
		Short: "Generates sql by excel",
		Run: func(cmd *cobra.Command, args []string) {
			// 获取当前文件夹所有文件信息 不包括文件夹内的文件
			filesInfo, err := resolveExcel.GetFiles()
			if err != nil {
				return
			} else {
				resolveExcel.GenerateSqlByExcel(filesInfo)
			}
		},
	}

	return generateSql
}

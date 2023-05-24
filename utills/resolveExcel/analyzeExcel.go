package resolveExcel

import (
	"autoSql-cobra/config"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

// AnalyzeExcel 分析excel
func AnalyzeExcel(name string, pat string) {
	fmt.Printf("Analyzing excel (%s) \n", name)
	//创建sql文件夹
	err := os.MkdirAll("./sqls", 0755)
	if err != nil {
		fmt.Println(err.Error())
	}

	//创建sql文件
	file, err := os.Create(fmt.Sprintf("./sqls/%s.sql", strings.Split(name, ".")[0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)

	//打开当前目录excle文件
	wd, err := os.Getwd()
	f, err := excelize.OpenFile(fmt.Sprintf("%s/%s", wd, pat))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(f)

	//读取excel文件内容
	rows, err := f.GetRows("Sheet1")

	for _, row := range rows {
		str := fmt.Sprintf(
			"insert into %s (%s) values ('%s')\n",
			config.Config.Schema,
			strings.Join(config.Config.Fields, ","),
			strings.Join(row[:len(config.Config.Fields)], "','"),
		)
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("Analyzing excel (%s) end \n", name)
}

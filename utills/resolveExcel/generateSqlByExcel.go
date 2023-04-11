package resolveExcel

import (
	"autoSql-cobra/config"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var wg sync.WaitGroup

// GenerateSqlByExcel 根据当前excel文件生成 sql 文件
func GenerateSqlByExcel(excelFiles []ExcelFiles) {
	// viper 加载config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(err.Error())
		}
	}

	var c config.SqlConfig
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("Serialization configuration failed.%s\n", err)
		return
	}

	fileNum := 0
	// 分析excel
	for _, file := range excelFiles {
		if !file.isDir && file.ext == ".xlsx" {
			fmt.Println(fmt.Sprintf("find excel file (%s)", file.name))
			fileNum++
			wg.Add(1)
			go func(name string, path string) {
				AnalyzeExcel(name, path)
				defer wg.Done()
			}(file.name, file.path)
		}
	}

	wg.Wait()

	fmt.Println(fmt.Sprintf("find excel file (%d)", fileNum))
}

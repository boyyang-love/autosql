package main

import (
	"autoSql-cobra/cmd"
	"autoSql-cobra/config"
	"autoSql-cobra/utills/loadEnv"
	"fmt"
)

func main() {
	// 加载环境变量
	if err := loadEnv.LoadEnv(); err != nil {
		fmt.Printf("Error loading environment files, (%s)", err.Error())
	}
	// 加载初始化config
	config.InitConfig()
	cmd.Execute()
}

package resolveExcel

import (
	"fmt"
	"os"
	"path/filepath"
)

type ExcelFiles struct {
	name  string
	ext   string
	isDir bool
	path  string
}

type FileInfos struct {
	info []ExcelFiles
}

// GetFiles 获取当前文件文件信息 不包括文件夹下的文件
func GetFiles() (excelFiles []ExcelFiles, err error) {
	err = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fmt.Println(fmt.Sprintf("正在分析文件----(%s)", info.Name()))
		excelFiles = append(
			excelFiles,
			ExcelFiles{
				name:  info.Name(),
				ext:   filepath.Ext(fmt.Sprintf("./%s", info.Name())),
				isDir: info.IsDir(),
				path:  path,
			},
		)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return excelFiles, nil
}

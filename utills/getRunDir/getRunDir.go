package getRunDir

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunDir() string {
	// 如果是开发环境 使用当前路径
	if os.Getenv("active") == "dev" {
		return "."
	}

	file, _ := exec.LookPath(os.Args[0])
	// 获取包含可执行文件名称的路径
	path, _ := filepath.Abs(file)
	// 获取可执行文件所在目录
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]

	return strings.Replace(ret, "\\", "/", -1)
}

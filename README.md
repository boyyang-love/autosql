### 根据excel生成sql文件(基于cobra的命令行工具)

#### 使用方法
1. 直接将bin目录文件下载后将文件添加到环境变量中
2. 通过 `autosql version`查看是否配置成功
3. 通过 `autosql schema -s [表名]`设置表名
4. 通过 `autosql fields -f [name,age,sex]`设置字段名称
5. 通过 `autosql auto`会自动检测当前文件夹的所有excel文件并读取相关内容生成sql文件夹

#### 技术栈
* go    v1.19
* cobra v1.6.1
* viper v1.15.0
* excelize v2.7.0
* godotenv v1.5.1
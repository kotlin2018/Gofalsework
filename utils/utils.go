package utils

import "os"

// 创建文件，并将内容写入文件
func WriteToFile(filename,content string){
	// 创建文件句柄
	file, err := os.Create(filename)
	if err !=nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil{
		panic(err)
	}
}

// 判断文件或者目录是否存在 (传入文件路径)
// bool = true 文件存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
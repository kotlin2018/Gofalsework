package main

import (
	"github.com/spf13/pflag"
	commands "Gofalsework/commands/new"
	"Gofalsework/utils"
	"log"
	"os"
)

func main() {
	//将命令行解析成标志
	//解析命令行(例如终端)输入的参数
	pflag.Parse()

	//获取当前目录对应的根目录的路径
	rootPath,_ := os.Getwd()

	//如果文件存在,输出提示信息
	if utils.FileIsExist(rootPath) {
		log.Printf("Application '%s' already exist",rootPath)
		os.Exit(0)
	}

	// 定义变量接收命令行参数
	appName := ""

	// 遍历命令行参数
	// 如果命令行没有输入参数、或者参数是 new
	//
	for v := range pflag.Args(){
		if v == 0{
			if pflag.Args()[v] == "new" {
				if len(pflag.Args()) >1 {
					//命令行有第二个参数
					if pflag.Args()[v+1] != ""{
						// 第二个参数就是将要创建的项目名称
						appName = pflag.Args()[v+1]
						// 执行创建项目的函数 ，在当前目录下创建项目
						commands.CreatedApi(rootPath,appName)
					}
				}else{
					goto App
				}
			} else if pflag.Args()[v] == "api" {
				if len(pflag.Args()) >1 {
					if pflag.Args()[v+1] != ""{
						appName = pflag.Args()[v+1]
						commands.CreatedWeb(rootPath,appName)
					}
				}else {
					goto App
				}
			}
		}
	App:
		//默认项目名
		appName = "irisTool"
		commands.CreatedWeb(rootPath,appName)
	}
}

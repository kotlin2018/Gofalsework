package commands

import (
	"Gofalsework/utils"
	"log"
	"os"
	"path"
	"strings"
)

var main = `
package main

import (
	//"github.com/kataras/iris/v12"
	//"github.com/spf13/pflag"
	//"github.com/spf13/viper"
	//
	//"{{.Appname}}/web/middleware"
	//"{{.Appname}}/config"
	//"{{.Appname}}/models"
	//"{{.Appname}}/router"
	"github.com/spf13/pflag"
	"log"
	"os"
	"path"
	"strings"
	"irisTool/utils"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	models.DB.Init()
	app := newApp()

	route.InitRouter(app)

	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	//	AllowCredentials: true,
	//	AllowedHeaders:   []string{"*"},
	//})
	//
	//app.Use(crs) //
	//app.StaticWeb("/assets", "./web/views/admin/assets")
	//app.RegisterView(iris.HTML("./web/views/admin", ".html"))
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.GetJWT().Serve)//是否启用jwt中间件
	app.Use(middleware.LogMiddle)//是否启用logrus中间件
	app.Configure(iris.WithOptimizations)

	return app
}
`




// 创建 api 的 目录结构
func CreatedApi(appPath string, appName string) {
	log.Println("Creating application...")
	log.Printf("当前所在路径为%s\n",appPath)
	//创建目录树 appName是项目名，即顶级目录
	os.MkdirAll(appName, 0755)
	os.Mkdir(path.Join(appName, "conf"), 0755)
	os.Mkdir(path.Join(appName, "controllers"), 0755)
	os.Mkdir(path.Join(appName, "models"), 0755)
	os.Mkdir(path.Join(appName, "routers"), 0755)
	os.Mkdir(path.Join(appName, "swagger"), 0755)
	os.Mkdir(path.Join(appName, "tests"), 0755)
	//os.Mkdir(path.Join(appName, "sysinit"), 0755)
	//os.Mkdir(path.Join(appName, "utils"), 0755)
	os.Mkdir(path.Join(appName, "common"), 0755)

	// 1. 将 config.yaml 放到 conf包（文件夹）下
	// 2. 加载conf的值 (conf 是conf.go文件中的一个 var变量)到config.yaml
	utils.WriteToFile(path.Join(appName, "conf", "config.yaml"), conf)
	//utils.WriteToFile(path.Join(appName, "config", "config.go"), config)
	utils.WriteToFile(path.Join(appName, "models", "init.go"), databaseInit)
	//utils.WriteToFile(path.Join(appName, "models", "log.go"), logger)
	utils.WriteToFile(path.Join(appName, "models", "user.go"), user)
	utils.WriteToFile(path.Join(appName, "/swagger", "index.html"), index_html)
	utils.WriteToFile(path.Join(appName, "/swagger", "oauth2-redirect.html"),oauth2_redirect_html )
	utils.WriteToFile(path.Join(appName, "/swagger", "swagger.json"),swagger_json)
	utils.WriteToFile(path.Join(appName, "/swagger", "swagger.yml"), swagger_yml)
	utils.WriteToFile(path.Join(appName, "/swagger", "swagger-ui.css"),swagger_ui_css )
	utils.WriteToFile(path.Join(appName, "/swagger", "swagger-ui.js"), swagger_ui_js)
	utils.WriteToFile(path.Join(appName, "/swagger", "swagger-ui-bundle.js"),swagger_ui_bundle_js )
	utils.WriteToFile(path.Join(appName, "/swagger", "swagger-ui-standalone-preset.js"),swagger_ui_standalone_preset_js )
	utils.WriteToFile(path.Join(appName, "routers", "router.go"), router)
	//utils.WriteToFile(path.Join(appName, "main.go"), main)
	
	// 在service 字符串中，把{{.Appname}} 替换为 appName
	//utils.WriteToFile(path.Join(appName, "service", "TestService.go"), strings.Replace(service, "{{.Appname}}", appName, -1))
	//utils.WriteToFile(path.Join(appName, "repositories", "TestRepo.go"), strings.Replace(repositories, "{{.Appname}}", appName, -1))

	//utils.WriteToFile(path.Join(appName, "/web/controllers", "TestController.go"), controllers)
	//utils.WriteToFile(path.Join(appName, "/web/controllers", "Common.go"), common)
	//utils.WriteToFile(path.Join(appName, "/web/middleware", "jwt.go"), strings.Replace(jwt, "{{.Appname}}", appName, -1))
	//utils.WriteToFile(path.Join(appName, "/web/middleware", "logrus.go"), strings.Replace(logrus, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "main.go"), strings.Replace(main, "{{.Appname}}", appName, -1))
	log.Println("new application successfully created!")
}

// 创建完整的 Web 目录结构
func CreatedWeb(appPath string, appName string) {
	log.Println("Creating application...")
	log.Printf("当前所在路径为%s\n",appPath)

	//创建目录树 appName是项目名，即顶级目录
	os.MkdirAll(appName, 0755)
	os.Mkdir(path.Join(appName, "conf"), 0755)
	//os.Mkdir(path.Join(appName, "controllers"), 0755)
	os.MkdirAll(path.Join(appName, "/api/controllers"), 0755)
	os.MkdirAll(path.Join(appName, "/api/models"), 0755)
	os.MkdirAll(path.Join(appName, "/api/routers"), 0755)
	//os.Mkdir(path.Join(appName, "static"), 0755)
	//os.Mkdir(path.Join(appName, "views"), 0755)

	os.MkdirAll(path.Join(appName, "/web/store"), 0755)
	os.MkdirAll(path.Join(appName, "/web/views"), 0755)
	//os.MkdirAll(path.Join(appName, "/web/static"), 0755)

	os.MkdirAll(path.Join(appName, "/web/static/css"), 0755)
	os.MkdirAll(path.Join(appName, "/web/static/img"), 0755)
	os.MkdirAll(path.Join(appName, "/web/static/js"), 0755)
	os.MkdirAll(path.Join(appName, "/web/static/vue"), 0755)

	os.MkdirAll(path.Join(appName, "/web/middleware"), 0755)
	os.Mkdir(path.Join(appName, "tests"), 0755)
	os.Mkdir(path.Join(appName, "common"), 0755)

	// 1. 将 config.yaml 放到 conf包（文件夹）下
	// 2. 加载conf的值 (conf 是conf.go文件中的一个 var变量)到config.yaml
	utils.WriteToFile(path.Join(appName, "conf", "config.yaml"), conf)
	utils.WriteToFile(path.Join(appName, "/api/controllers", "default.go"), handler) //handler是变量名
	utils.WriteToFile(path.Join(appName, "/api/models", "init.go"), databaseInit)
	utils.WriteToFile(path.Join(appName, "/api/models", "log.go"), logger)
	utils.WriteToFile(path.Join(appName, "/api/routers", "router.go"), router) //rout 是var 变量

	utils.WriteToFile(path.Join(appName, "/web/store", "index.js"), store)
	utils.WriteToFile(path.Join(appName, "/web/views", "index.tpl"), html) // common 里要写html

	utils.WriteToFile(path.Join(appName, "/web/static/css", "index.css"), css) // css 里要写css
	utils.WriteToFile(path.Join(appName, "/web/static/js", "index.js"), js)
	utils.WriteToFile(path.Join(appName, "/web/static/vue", "App.vue"), vue)

	utils.WriteToFile(path.Join(appName, "/web/middleware", "jwt.go"), strings.Replace(jwt, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "/web/middleware", "logrus.go"), strings.Replace(logrus, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "main.go"), strings.Replace(main, "{{.Appname}}", appName, -1))
	log.Println("new application successfully created!")
}
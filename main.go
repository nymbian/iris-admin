package main

import (
	"io"
	"github.com/nymbian/iris-admin/common"
	"github.com/nymbian/iris-admin/libs"
	"github.com/nymbian/iris-admin/model"
	"github.com/nymbian/iris-admin/route"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {
	os.Mkdir("logs", 0755)
	logFile, err := os.OpenFile("./logs/run_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln("open log file failed", err)
	}

	//日志
	LogInfo = log.New(io.MultiWriter(logFile), "【Info】:", log.Ldate|log.Ltime|log.Lshortfile)   //LogInfo.Println(1, 2, 3)
	LogError = log.New(io.MultiWriter(logFile), "【Error】:", log.Ldate|log.Ltime|log.Lshortfile) //LogError.Println(4, 5, 6)

	config.AddConfigPath("./configs")
	config.SetConfigName("mysql")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbConfig := libs.DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
	}
	libs.DB = dbConfig.InitDB()
	if config.GetBool("default.sql_log") {
		libs.DB.LogMode(true)
	}
}

func main() {
	app := iris.New()
	config.SetConfigName("app")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误, %s", err)
	}
	tmpl := iris.HTML("./views", ".html").Layout(config.GetString("site.DefaultLayout"))
	if config.GetBool("site.APPDebug") == true {
		app.Logger().SetLevel("debug") //设置debug
		tmpl.Reload(true)
	}

	tmpl.AddFunc("TimeToDate", libs.TimeToDate)
	tmpl.AddFunc("strToHtml", libs.StrToHtml)

	app.RegisterView(tmpl)
	app.Favicon("./favicon.ico")
	//app.Use(iris.Gzip)

	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())

	app.HandleDir("/static", "./static")   //设置静态文件目录
	app.HandleDir("/uploads", "./uploads") //设置静态文件目录

	//设置公共页面输出
	app.Use(func(ctx iris.Context) {
		if auth := common.SessManager.Start(ctx).Get("admin_user"); auth != nil {
			admin_user, _ := auth.(map[string]interface{})
			var admin_model model.Admin
			admin_id, _ := admin_user["id"].(uint)
			adminInfo, _ := admin_model.AdminInfo(admin_id)
			if adminInfo.Avatar == "" {
				adminInfo.Avatar = "/static/adminlit/dist/img/user2-160x160.jpg"
			}
			ctx.ViewData("adminInfo", adminInfo)
		}
		ctx.ViewData("Title", config.GetString("site.DefaultTitle"))
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.Next()
	})

	//设置错误模版
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.HTML("<center>很抱歉！当前页面错误,错误代码:" + strconv.Itoa(ctx.GetStatusCode()) + "</center>")
	})

	route.Routes(app)

	//应用配置文件
	app.Configure(iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

	//Run
	www := app.Party("www.")
	{
		currentRoutes := app.GetRoutes()
		for _, r := range currentRoutes {
			www.Handle(r.Method, r.Tmpl().Src, r.Handlers...)
		}
	}
	app.Run(iris.Addr(config.GetString("server.domain") + ":" + config.GetString("server.port")))
}

package route

import (
	commons "iris-admin/common"
	controllers "iris-admin/controller"
	"iris-admin/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Routes(app *iris.Application) {
	//首页路由
	mvc.New(app.Party("/")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.IndexController))

	//登录路由
	mvc.New(app.Party("/login")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.LoginController))

	//系统路由
	mvc.New(app.Party("/system", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.SystemController))
	//管理员管理
	mvc.New(app.Party("/administrator", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.AdministratorController))
	//分类管理
	mvc.New(app.Party("/category", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.CategoryController))
	//菜单管理
	mvc.New(app.Party("/menu", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.MenuController))
	//内容管理
	mvc.New(app.Party("/news", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.NewsController))
}

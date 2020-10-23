package route

import (
	"github.com/nymbian/iris-admin/common"
	controllers "github.com/nymbian/iris-admin/controller"
	"github.com/nymbian/iris-admin/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Routes(app *iris.Application) {
	//首页路由
	mvc.New(app.Party("/")).
		Register(common.SessManager.Start).
		Handle(new(controllers.IndexController))

	//登录路由
	mvc.New(app.Party("/login")).
		Register(common.SessManager.Start).
		Handle(new(controllers.LoginController))

	//系统路由
	mvc.New(app.Party("/system", middleware.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.SystemController))
	//管理员管理
	mvc.New(app.Party("/administrator", middleware.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.AdministratorController))
	//分类管理
	mvc.New(app.Party("/category", middleware.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.CategoryController))
	//菜单管理
	mvc.New(app.Party("/menu", middleware.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.MenuController))
	//内容管理
	mvc.New(app.Party("/news", middleware.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.NewsController))
}

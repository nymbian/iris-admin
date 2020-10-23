package controllers

import (
	"github.com/nymbian/iris-admin/common"
	"github.com/nymbian/iris-admin/model"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

var MenuModel = model.Menu{}

type MenuController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *MenuController) Get() mvc.View {
	Menu := model.Menu{}
	list := Menu.List()
	model.MenuListTree = []model.Menu{}
	list = Menu.GetTree(list, 0, 0)
	return mvc.View{
		Name: "menu/list.html",
		Data: iris.Map{
			"Title": "菜单列表",
			"list":  list,
		},
	}
}

func (c *MenuController) GetAddMenu() mvc.View {
	Menu := model.Menu{}
	list := Menu.List()
	model.MenuListTree = []model.Menu{}
	list = Menu.GetTree(list, 0, 0)
	return mvc.View{
		Name: "menu/add.html",
		Data: iris.Map{
			"Title": "新增菜单",
			"list":  list,
		},
	}
}

func (c *MenuController) PostAddMenu() {
	if err := MenuModel.MenuAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/menu")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *MenuController) GetUpdateMenuBy(id uint) mvc.View {
	MenuInfo, err := MenuModel.MenuInfo(id)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	Menu := model.Menu{}
	list := Menu.List()
	model.MenuListTree = []model.Menu{}
	list = Menu.GetTree(list, 0, 0)

	return mvc.View{
		Name: "menu/update.html",
		Data: iris.Map{
			"Title":          "菜单修改",
			"UpdateMenuInfo": MenuInfo,
			"list":           list,
		},
	}
}

func (c *MenuController) PostUpdateMenu() {
	if err := MenuModel.MenuUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/menu")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *MenuController) GetDelMenuBy(id uint) {
	if err := MenuModel.MenuDel(id); err == nil {
		c.Ctx.Redirect("/menu")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

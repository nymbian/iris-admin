package controllers

import (
	"github.com/nymbian/iris-admin/common"
	"github.com/nymbian/iris-admin/model"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

var CategoryModel = model.Category{}

type CategoryController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *CategoryController) Get() mvc.View {
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)
	return mvc.View{
		Name: "category/list.html",
		Data: iris.Map{
			"Title": "分类列表",
			"list":  list,
		},
	}
}

func (c *CategoryController) GetAddCategory() mvc.View {
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)
	return mvc.View{
		Name: "category/add.html",
		Data: iris.Map{
			"Title": "新增分类",
			"list":  list,
		},
	}
}

func (c *CategoryController) PostAddCategory() {
	if err := CategoryModel.CategoryAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/category")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *CategoryController) GetUpdateCategoryBy(id uint) mvc.View {
	categoryInfo, err := CategoryModel.CategoryInfo(id)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)

	return mvc.View{
		Name: "category/update.html",
		Data: iris.Map{
			"Title":              "分类修改",
			"UpdateCategoryInfo": categoryInfo,
			"list":               list,
		},
	}
}

func (c *CategoryController) PostUpdateCategory() {
	if err := CategoryModel.CategoryUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/category")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *CategoryController) GetDelCategoryBy(id uint) {
	if err := CategoryModel.CategoryDel(id); err == nil {
		c.Ctx.Redirect("/category")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

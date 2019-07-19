package controllers

import (
	"html"
	commons "iris-admin/common"
	"iris-admin/libs"
	"iris-admin/model"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type AdministratorController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *AdministratorController) Get() mvc.View {
	Admin := model.Admin{}
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := Admin.List(page)
	return mvc.View{
		Name: "administrator/list.html",
		Data: iris.Map{
			"Title":    "管理员列表",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}

func (c *AdministratorController) GetUpdateAdminBy(id uint) mvc.View {
	adminInfo, err := admin_model.AdminInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}

	if adminInfo.Avatar == "" {
		adminInfo.Avatar = "/static/images/NoteFound.jpg"
	}
	return mvc.View{
		Name: "administrator/update.html",
		Data: iris.Map{
			"Title":           "资料修改",
			"UpdateAdminInfo": adminInfo,
		},
	}
}

func (c *AdministratorController) PostUpdateAdmin() {
	
	var postValues map[string][]string
	postValues = c.Ctx.FormValues()
	admin_id := postValues["id"][0]
	int_admin_id, _ := strconv.Atoi(admin_id)
	err, filePath := libs.UploadFile("avatar", c.Ctx)
	if err == false {
		//commons.DefaultErrorShow(filePath, c.Ctx)
		//return
		filePath = postValues["avatar_old"][0]
	}
	delete(postValues, "id")
	delete(postValues, "avatar_old")
	if err := admin_model.AdminUpdate(postValues, uint(int_admin_id), filePath); err == nil {
		c.Ctx.Redirect("/administrator/update/admin/" + admin_id)
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *AdministratorController) GetAddAdmin() mvc.View {
	return mvc.View{
		Name: "administrator/add.html",
		Data: iris.Map{
			"Title": "新增管理员",
		},
	}
}

func (c *AdministratorController) PostAddAdmin() {
	err, filePath := libs.UploadFile("avatar", c.Ctx)
	if err == false {
		commons.DefaultErrorShow(filePath, c.Ctx)
		return
	}

	if err := admin_model.AddUpdate(c.Ctx.FormValues(), filePath); err == nil {
		c.Ctx.Redirect("/administrator")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *AdministratorController) GetUpdatePasswordBy(id uint) mvc.View {
	adminInfo, err := admin_model.AdminInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "administrator/password.html",
		Data: iris.Map{
			"Title":   "密码修改",
			"Id":      id,
			"Account": adminInfo.Account,
		},
	}
}

func (c *AdministratorController) PostUpdatePassword() {
	id := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("id")))
	password := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	Repassword := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("Repassword")))
	int_admin_id, _ := strconv.Atoi(id)
	if err := admin_model.AdminPasswodUpdate(uint(int_admin_id), password, Repassword); err == nil {
		c.Ctx.Redirect("/administrator")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *AdministratorController) GetDelAdminBy(id uint) {
	if err := admin_model.AdminDel(id); err == nil {
		c.Ctx.Redirect("/administrator")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

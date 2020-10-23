package middleware

import (
	"github.com/nymbian/iris-admin/common"
	"github.com/kataras/iris/context"
)

func SessionLoginAuth(Ctx *context.Context) {
	if auth := common.SessManager.Start(Ctx).Get("admin_user"); auth == nil {
		Ctx.Redirect("/login")
		return
	}
	Ctx.Next()
}
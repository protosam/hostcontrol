package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
)

func set_error(msg string, ctx *macaron.Context) {
	ctx.SetCookie("err_str", msg)
}

func set_info(msg string, ctx *macaron.Context) {
	ctx.SetCookie("info_str", msg)
}

func header(ctx *macaron.Context) (string) {
	var tpl vision.New
        tpl.TemplateFile("template/overall.tpl")

	user, auth := util.Auth(ctx, "any")
	if auth {
		tpl.Assign("username", user.System_username)
	}

        tpl.Parse("header")


	err_str := ctx.GetCookie("err_str")
	if err_str != "" {
		tpl.Assign("message", err_str)
		tpl.Parse("header/error")
		ctx.SetCookie("err_str", "")
	}

	info_str := ctx.GetCookie("info_str")
	if info_str != "" {
		tpl.Assign("message", info_str)
		tpl.Parse("header/info")
		ctx.SetCookie("info_str", "")
	}

	return tpl.Out()
}


func footer(ctx *macaron.Context) (string) {
	var tpl vision.New
        tpl.TemplateFile("template/overall.tpl")
        tpl.Parse("footer")
        return tpl.Out()
}


func die(ctx *macaron.Context, msg string) (string) {
	var tpl vision.New
        tpl.TemplateFile("template/error.tpl")
        tpl.Assign("message", msg)
        tpl.Parse("error")
        return header(ctx) + tpl.Out() + footer(ctx)

}


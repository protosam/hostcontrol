package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"strings"
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

	hcuser, auth := util.Auth(ctx, "any")
	if auth {
		tpl.Assign("username", hcuser.System_username)
	}

	hostname := string(ctx.Req.Header.Get("X-FORWARDED-HOST"))
	if hostname == "" {
		hostname = string(ctx.Req.Host)
	}
	hostname = strings.Split(hostname, ":")[0]

	tpl.Assign("console_url", "https://"+hostname+"/shellinabox")
    tpl.Parse("header")


    if auth {
        if (strings.Contains(hcuser.Privileges, "websites") || strings.Contains(hcuser.Privileges, "all")) && hcuser.System_username != "root" {
            tpl.Parse("header/websitesbtn");
        }
        if strings.Contains(hcuser.Privileges, "databases") || strings.Contains(hcuser.Privileges, "all") {
            tpl.Parse("header/databasesbtn");
        }
        if strings.Contains(hcuser.Privileges, "dns") || strings.Contains(hcuser.Privileges, "all") {
            tpl.Parse("header/dnsbtn");
        }
        if (strings.Contains(hcuser.Privileges, "mail") || strings.Contains(hcuser.Privileges, "all")) && hcuser.System_username != "root" {
            tpl.Parse("header/mailbtn");
        }
        if (strings.Contains(hcuser.Privileges, "ftpusers") || strings.Contains(hcuser.Privileges, "all")) && hcuser.System_username != "root" {
            tpl.Parse("header/ftpusersbtn");
        }
        if strings.Contains(hcuser.Privileges, "all") {
            tpl.Parse("header/firewallbtn");
        }
        if strings.Contains(hcuser.Privileges, "all") {
            tpl.Parse("header/servicesbtn");
        }
        if strings.Contains(hcuser.Privileges, "sysusers") || strings.Contains(hcuser.Privileges, "all") {
            tpl.Parse("header/usersbtn");
        }
    }


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


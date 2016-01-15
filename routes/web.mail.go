package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
	"strings"
)


func init(){
	route("/mail", mail, "BOTH")
	route("/mail/domain/add", mailadddomain, "BOTH")
	route("/mail/domain/delete", maildeletedomain, "BOTH")
	route("/mail/users/add", mailadduser, "BOTH")
	route("/mail/users/edit", mailedituser, "BOTH")
	route("/mail/users/delete", maildeleteuser, "BOTH")
}


func maildeleteuser(ctx *macaron.Context) (string){
    status := API("/api/mail/users/delete", ctx)
    
    email := util.Query(ctx, "email")
    
    if status == "success" {
        set_error("Deleted " + email + " successfully!", ctx)
        ctx.Redirect("/mail", 302)
        return "did it!"
    }
    
    set_error("Failed to delete " + email + ". Error given: " + status, ctx)
    ctx.Redirect("/mail", 302)
    
    return "Failed to delete user. Error given: " + status
}
func mailedituser(ctx *macaron.Context) (string){
    status := API("/api/mail/users/edit", ctx)
    
    email := util.Query(ctx, "email")
    
    if status == "success" {
        set_error("Updated " + email + " successfully!", ctx)
        ctx.Redirect("/mail", 302)
        return "did it!"
    }
    
    set_error("Failed to update " + email + ". Error given: " + status, ctx)
    ctx.Redirect("/mail", 302)
    
    return "Failed to update user. Error given: " + status
}

func mailadduser(ctx *macaron.Context) (string){
    status := API("/api/mail/users/add", ctx)
    
    username := util.Query(ctx, "username")
    domain := util.Query(ctx, "domain")
    
    if status == "success" {
        set_error("Added " + username + "@" + domain + " successfully!", ctx)
        ctx.Redirect("/mail", 302)
        return "did it!"
    }
    
    set_error("Failed to add user. Error given: " + status, ctx)
    ctx.Redirect("/mail", 302)
    
    return "Failed to add user. Error given: " + status
}

func maildeletedomain(ctx *macaron.Context) (string){
    status := API("/api/mail/domain/delete", ctx)
    
    domainname := util.Query(ctx, "domain")
    
    if status == "success" {
        set_error("Deleted " + domainname + " successfully!", ctx)
        ctx.Redirect("/mail", 302)
        return "did it!"
    }
    
    set_error("Failed to delete domain. Error given: " + status, ctx)
    ctx.Redirect("/mail", 302)
    
    return "Failed to delete domain. Error given: " + status
}


func mailadddomain(ctx *macaron.Context) (string){
    status := API("/api/mail/domain/add", ctx)
    
    domainname := util.Query(ctx, "domain")
    
    if status == "success" {
        set_error("Added " + domainname + " successfully!", ctx)
        ctx.Redirect("/mail", 302)
        return "did it!"
    }
    
    set_error("Failed to add domain. Error given: " + status, ctx)
    ctx.Redirect("/mail", 302)
    
    return "Failed to add domain. Error given: " + status
}

func mail(ctx *macaron.Context) (string){
	_, auth := util.Auth(ctx, "mail")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/mail.tpl")


	hostname := string(ctx.Req.Header.Get("X-FORWARDED-HOST"))
	if hostname == "" {
		hostname = string(ctx.Req.Host)
	}
	hostname = strings.Split(hostname, ":")[0]

	tpl.Assign("webmail_url", "https://"+hostname+"/roundcubemail")


	tpl.Parse("mail")

	// list domains and records
	dns_data := API("/api/mail/list", ctx)

	
	// map[domain]map[record_id]map[key]value
	data := make(map[string]map[string]map[string]string)
	json.Unmarshal([]byte(dns_data), &data)
	
	for domain, email_accounts := range data {
        tpl.Assign("domain_name", domain)
	    tpl.Parse("mail/domain")
	    
	    for key, email := range email_accounts {
	        if key == "placebo" {
	            continue
	        }
            tpl.Assign("email", email["email"])
            tpl.Assign("email_id", email["email_id"])
    	    tpl.Parse("mail/domain/email")
	    }
	}

	return header(ctx) + tpl.Out() + footer(ctx)
}


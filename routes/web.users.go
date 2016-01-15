package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	
	"encoding/json"
	"strings"
)


func init(){
	route("/users", users, "BOTH")
	route("/users/add", useradd, "POST")
	route("/users/delete", userdelete, "BOTH")
	route("/users/sudo", sudo, "BOTH")
}


func users(ctx *macaron.Context) (string){
	hcuser, auth := util.Auth(ctx, "sysusers")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}

	var tpl vision.New
	tpl.TemplateFile("template/users.tpl")
	tpl.Parse("users")
	
	if strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_all")
	}
	if strings.Contains(hcuser.Privileges, "websites") || strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_websites")
	}
	if strings.Contains(hcuser.Privileges, "mail") || strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_mail")
	}
	if strings.Contains(hcuser.Privileges, "databases") || strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_databases")
	}
	if strings.Contains(hcuser.Privileges, "ftpusers") || strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_ftpusers")
	}
	if strings.Contains(hcuser.Privileges, "dns") || strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_dns")
	}
	if strings.Contains(hcuser.Privileges, "sysusers") || strings.Contains(hcuser.Privileges, "all") {
    	tpl.Parse("users/perms_sysusers")
	}

	userdata := API("/api/users/list", ctx)
	
	users := make(map[string]map[string]string)
	json.Unmarshal([]byte(userdata), &users)
	
	for _, user := range users {
        tpl.Assign("hostcontrol_id", user["hostcontrol_id"])
        tpl.Assign("system_username", user["system_username"])
        tpl.Assign("privileges", user["privileges"])
        tpl.Assign("owned_by", user["owned_by"])
        tpl.Assign("login_token", user["login_token"])
        tpl.Assign("email_address", user["email_address"])
        
	    tpl.Parse("users/user")
	}
	
	
	return header(ctx) + tpl.Out() + footer(ctx)
}


func useradd(ctx *macaron.Context) (string) {
    status := API("/api/users/add", ctx)
    
    username := util.Query(ctx, "username")
    
    if status == "success" {
        set_error("Added " + username + " successfully!", ctx)
        ctx.Redirect("/users", 302)
        return "did it!"
    }
    
    set_error("Failed to add user. Error given: " + status, ctx)
    ctx.Redirect("/users", 302)
    
    return "Failed to add user. Error given: " + status
}




func userdelete(ctx *macaron.Context) (string) {
    status := API("/api/users/delete", ctx)
    
    username := util.Query(ctx, "username")
    
    if status == "success" {
        set_error("Deleted " + username + " successfully!", ctx)
        ctx.Redirect("/users", 302)
        return "did it!"
    }
    
    set_error("Failed to delete user. Error given: " + status, ctx)
    ctx.Redirect("/users", 302)
    
    return "Failed to add user. Error given: " + status
}



func sudo(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "sysusers")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}
    
    username := util.Query(ctx, "username")
    
    if ! util.ChkPaternity(hcuser.System_username, username) {
        set_error("Failed to sudo to "+username+"!", ctx)
        ctx.Redirect("/users", 302)
        return "failed!"
        
    }
    
    ctx.SetCookie("sudo", username, 864000)
    set_error("You are now logged in as " + username + "! Clicking logout will switch back to " + hcuser.System_username + ".", ctx)
    ctx.Redirect("/dashboard", 302)
    return "success"
}


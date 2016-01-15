package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
)


func init(){
	route("/ftpusers", ftpusers, "BOTH")
	route("/ftpusers/add", ftpuseradd, "BOTH")
	route("/ftpusers/delete", ftpuserdelete, "BOTH")
	route("/ftpusers/edit", ftpuseredit, "BOTH")
}


func ftpusers(ctx *macaron.Context) (string){
	hcuser, auth := util.Auth(ctx, "ftpusers")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/ftpusers.tpl")
	
	tpl.Assign("homedir", hcuser.HomeDir)
	tpl.Parse("ftpusers")
	
	userdata := API("/api/ftpusers/list", ctx)

	users := make(map[string]map[string]string)
	json.Unmarshal([]byte(userdata), &users)
	
	for _, user := range users {
        tpl.Assign("username", user["username"])
        tpl.Assign("homedir", user["homedir"])
        
	    tpl.Parse("ftpusers/user")
	}

	return header(ctx) + tpl.Out() + footer(ctx)
}


func ftpuseradd(ctx *macaron.Context) (string) {
    status := API("/api/ftpusers/add", ctx)
    
    username := util.Query(ctx, "ftpuser")
    
    if status == "success" {
        set_error("Added " + username + " successfully!", ctx)
        ctx.Redirect("/ftpusers", 302)
        return "did it!"
    }
    
    set_error("Failed to add user. Error given: " + status, ctx)
    ctx.Redirect("/ftpusers", 302)
    
    return "Failed to add user. Error given: " + status
}


func ftpuseredit(ctx *macaron.Context) (string) {
    status := API("/api/ftpusers/edit", ctx)
    
    username := util.Query(ctx, "ftpuser")
    
    if status == "success" {
        set_error("Updated " + username + " successfully!", ctx)
        ctx.Redirect("/ftpusers", 302)
        return "did it!"
    }
    
    set_error("Failed to update user. Error given: " + status, ctx)
    ctx.Redirect("/ftpusers", 302)
    
    return "Failed to update user. Error given: " + status
}




func ftpuserdelete(ctx *macaron.Context) (string) {
    status := API("/api/ftpusers/delete", ctx)
    
    username := util.Query(ctx, "ftpuser")
    
    if status == "success" {
        set_error("Deleted " + username + " successfully!", ctx)
        ctx.Redirect("/ftpusers", 302)
        return "did it!"
    }
    
    set_error("Failed to delete user. Error given: " + status, ctx)
    ctx.Redirect("/ftpusers", 302)
    
    return "Failed to add user. Error given: " + status
}
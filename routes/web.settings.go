package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
)


func init(){
	route("/settings", settings, "BOTH")
    route("/settings/update", updatesettings, "BOTH")
    route("/settings/tokens/add", addtoken, "BOTH")
    route("/settings/tokens/delete", deletetoken, "BOTH")
}


func deletetoken(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}
    
    token := util.Query(ctx, "token")
    
    
    db, _ := util.MySQL();
    defer db.Close()
    

	ustmt, _ := db.Prepare("DELETE FROM `hostcontrol`.`hostcontrol_user_tokens` WHERE `token`=? and hostcontrol_id=?")
	ustmt.Exec(token, hcuser.Hostcontrol_id)
	ustmt.Close()
    
    set_error("Token deleted.", ctx)
    ctx.Redirect("/settings", 302)
    
    return ""
}


func addtoken(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}
    
    description := util.Query(ctx, "description")
    token := util.MkToken()
    
    
    db, _ := util.MySQL();
    defer db.Close()
    
	xstmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`hostcontrol_user_tokens` set `token`=?, `hostcontrol_id`=?, `description`=?, token_id=null")
	_, err := xstmt.Exec(token, hcuser.Hostcontrol_id, description)
	xstmt.Close()

	if err != nil {
        set_error("Failed to create new token.", ctx)
        ctx.Redirect("/settings", 302)
	    return "Failed to create new token."
	}
    
    set_error("Created new token.", ctx)
    ctx.Redirect("/settings", 302)
    
    return ""
}

func updatesettings(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}
	
    password := util.Query(ctx, "password")
    new_password := util.Query(ctx, "new_password")
    new_password_verify := util.Query(ctx, "new_password_verify")
    
    if password == "" || new_password == "" || new_password_verify == "" {
        set_error("Failed to update settings. Error given: missing a password field", ctx)
        ctx.Redirect("/settings", 302)
        return ""
    }
    
    if new_password != new_password_verify {
        set_error("Failed to update settings. Error given: new passwords don't match", ctx)
        ctx.Redirect("/settings", 302)
        return ""
    }
    
    if ! chklogin(hcuser.System_username, password) {
        set_error("Failed to update settings. Error given: current password incorrect", ctx)
        ctx.Redirect("/settings", 302)
        return ""
    }
    
    chpassword(hcuser.System_username, new_password)


    set_error("Settings updated successfully.", ctx)
    ctx.Redirect("/settings", 302)
    
    return ""
}

func settings(ctx *macaron.Context) (string){
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/settings.tpl")
	tpl.Assign("username", hcuser.System_username)
	tpl.Parse("settings")


    db, _ := util.MySQL();
    defer db.Close()

    // tokens
    stmt, _ := db.Prepare("select * from hostcontrol_user_tokens where hostcontrol_id=?")
    rows, _ := stmt.Query(hcuser.Hostcontrol_id)
    stmt.Close()
    

    for rows.Next() {
        var token string
        var hostcontrol_id string
        var description string
        var token_id string
        
        rows.Scan(&token, &hostcontrol_id, &description, &token_id)
        
        tpl.Assign("raw_token", token)
        tpl.Assign("token", hostcontrol_id + "/" + token)
        tpl.Assign("description", description)
        
        tpl.Parse("settings/token")
    }


	return header(ctx) + tpl.Out() + footer(ctx)
}


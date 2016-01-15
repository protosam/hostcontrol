package routes

import (
	"strings"
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
)


func init(){
	route("/databases", databases, "BOTH")
	route("/databases/grants/add", databasegrantadd, "BOTH")
	route("/databases/grants/delete", databasegrantdelete, "BOTH")
	route("/databases/users/edit", databaseusersedit, "BOTH")
	route("/databases/users/delete", databaseusersdelete, "BOTH")
	route("/databases/users/add", databaseusersadd, "BOTH")
	route("/databases/delete", databasedelete, "BOTH")
	route("/databases/add", databaseadd, "BOTH")
}

func databaseusersadd(ctx *macaron.Context) (string){
    status := API("/api/sql/users/add", ctx)
    
    db_user := util.Query(ctx, "db_user")
    
    if status == "success" {
        set_error("Created " + db_user + " successfully!", ctx)
        ctx.Redirect("/databases", 302)
        return "Created " + db_user + " successfully!"
    }
    
    set_error("Failed to create " + db_user + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to create " + db_user + "! Error given: " + status
}

func databaseadd(ctx *macaron.Context) (string){
    status := API("/api/sql/databases/add", ctx)
    
    db_name := util.Query(ctx, "db_name")
    
    if status == "success" {
        set_error("Created " + db_name + " successfully!", ctx)
        ctx.Redirect("/databases", 302)
        return "Created " + db_name + " successfully!"
    }
    
    set_error("Failed to create " + db_name + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to create " + db_name + "! Error given: " + status
}

func databasedelete(ctx *macaron.Context) (string){
    status := API("/api/sql/databases/delete", ctx)
    
    db_name := util.Query(ctx, "db_name")
    
    if status == "success" {
        set_error("Deleted " + db_name + " successfully!", ctx)
        ctx.Redirect("/databases", 302)
        return "Deleted " + db_name + " successfully!"
    }
    
    set_error("Failed to delete " + db_name + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to delete " + db_name + "! Error given: " + status
}

func databaseusersdelete(ctx *macaron.Context) (string){
    status := API("/api/sql/users/delete", ctx)
    
    db_user := util.Query(ctx, "db_user")
    
    if status == "success" {
        set_error("Deleted " + db_user + " successfully!", ctx)
        ctx.Redirect("/databases", 302)
        return "Deleted " + db_user + " successfully!"
    }
    
    set_error("Failed to delete " + db_user + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to delete " + db_user + "! Error given: " + status
}

func databaseusersedit(ctx *macaron.Context) (string){
    status := API("/api/sql/users/edit", ctx)
    
    db_user := util.Query(ctx, "db_user")
    
    if status == "success" {
        set_error("Modified " + db_user + " successfully!", ctx)
        ctx.Redirect("/databases", 302)
        return "Modified " + db_user + " successfully!"
    }
    
    set_error("Failed to update " + db_user + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to update " + db_user + "! Error given: " + status
}

func databasegrantadd(ctx *macaron.Context) (string){
    status := API("/api/sql/grants/add", ctx)
    
    db_user := util.Query(ctx, "db_user")
    db_name := util.Query(ctx, "db_name")
    
    if status == "success" {
        set_error("Added " + db_user + " to database " + db_name + " successfully!", ctx)
        ctx.Redirect("/databases", 302)
        return "Added " + db_user + " to database " + db_name + " successfully!"
    }
    
    set_error("Failed to add " + db_user + " to database " + db_name + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to add " + db_user + " to database " + db_name + "! Error given: " + status
}

func databasegrantdelete(ctx *macaron.Context) (string){
    status := API("/api/sql/grants/delete", ctx)
    
    db_user := util.Query(ctx, "db_user")
    db_name := util.Query(ctx, "db_name")
    
    if status == "success" {
        set_error("Removed " + db_user + " from database " + db_name + "!", ctx)
        ctx.Redirect("/databases", 302)
        return "Removed " + db_user + " from database " + db_name + "!"
    }
    
    set_error("Failed to remove " + db_user + " from database " + db_name + "! Error given: " + status, ctx)
    ctx.Redirect("/databases", 302)
    
    return "Failed to remove " + db_user + " from database " + db_name + "! Error given: " + status
}

func databases(ctx *macaron.Context) (string){
	_, auth := util.Auth(ctx, "databases")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}

	//hostname := string(ctx.Req.Host)
	hostname := string(ctx.Req.Header.Get("X-FORWARDED-HOST"))
	if hostname == "" {
		hostname = string(ctx.Req.Host)
	}
	hostname = strings.Split(hostname, ":")[0]

	var tpl vision.New

	tpl.TemplateFile("template/databases.tpl")

	tpl.Assign("phpmyadmin_url", "https://"+hostname+"/phpmyadmin")
	tpl.Parse("databases")
	

	// list db users
	dbuser_data := API("/api/sql/users/list", ctx)
	var users []string
	json.Unmarshal([]byte(dbuser_data), &users)
	
	for _, user := range users {
        tpl.Assign("db_user", user)
        
	    tpl.Parse("databases/user")
	}
	// end: list db users
	
	// list databases
	db_data := API("/api/sql/databases/list", ctx)
	var databases []string
	json.Unmarshal([]byte(db_data), &databases)
	
	for _, db_name := range databases {
	    
	    // list grants
	    grant_data := API("/api/sql/grants/list?db_name="+db_name, ctx)
    	var grants []string
    	json.Unmarshal([]byte(grant_data), &grants)
    	
    	
        tpl.Assign("db_name", db_name)
        
	    tpl.Parse("databases/database")
	    
	    // grant add dropdown
    	for _, user := range users {
            tpl.Assign("db_user", user)
            tpl.Assign("db_name", db_name)
    	    tpl.Parse("databases/database/add_grant")
    	}
    	// END: grant add dropdown
    
    	for _, gdb_user := range grants {
            tpl.Assign("db_user", gdb_user)
            tpl.Assign("db_name", db_name)
    	    tpl.Parse("databases/database/grant")
    	}
	    // end: list grants
	    
	}
	// end: list databases
	

	return header(ctx) + tpl.Out() + footer(ctx)
}


package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
	"strings"
)


func init(){
	route("/websites", websites, "BOTH")
	route("/websites/add", addsite, "POST")
	route("/websites/delete",deletesite, "BOTH")
	route("/websites/sslmanager", sslmanager, "BOTH")
	route("/websites/sslmanager/update", sslupdate, "BOTH")
}

func sslupdate(ctx *macaron.Context) (string){
    status := API("/api/web/domain/sslmanage", ctx)
    
    vhost_id := util.Query(ctx, "vhost_id")
    
    if status == "success" {
        set_error("Updated SSL settings successfully!", ctx)
        ctx.Redirect("/websites/sslmanager?vhost_id="+vhost_id, 302)
        return "did it!"
    }
    
    set_error("Failed to add domain. Error given: " + status, ctx)
    ctx.Redirect("/websites/sslmanager?vhost_id="+vhost_id, 302)
    
    return "Failed to update SSL for domain. Error given: " + status
}



func sslmanager(ctx *macaron.Context) (string){
	_, auth := util.Auth(ctx, "websites")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


    vhost_id := util.Query(ctx, "vhost_id")
    
	var tpl vision.New
	tpl.TemplateFile("template/websites.sslmanager.tpl")
	

	websites := API("/api/web/domain/list", ctx)
	
	domains := make(map[string]map[string]string)
	json.Unmarshal([]byte(websites), &domains)
	
	found := false
	for _, domain := range domains {
	    if domain["vhost_id"] == vhost_id {
            tpl.Assign("vhost_id", domain["vhost_id"])
            tpl.Assign("system_username", domain["system_username"])
            tpl.Assign("domain", domain["domain"])
            tpl.Assign("documentroot", domain["documentroot"])
            tpl.Assign("ipaddr", domain["ipaddr"])
            tpl.Assign("ssl_certificate", domain["ssl_certificate"])
            tpl.Assign("ssl_key", domain["ssl_key"])
            tpl.Assign("ssl_ca_certificate", domain["ssl_ca_certificate"])
            
            if domain["ssl_enabled"] == "Y" {
                tpl.Assign("ssl_enabled", "checked")
            }else{
                tpl.Assign("ssl_enabled", "")
            }
            found = true
    	}
	}
	
	if ! found {
        set_error("Failed to find requested domain.", ctx)
        ctx.Redirect("/websites", 302)
	    return ""
	}
	tpl.Parse("sslmanager")
	
	return header(ctx) + tpl.Out() + footer(ctx)
}

func deletesite(ctx *macaron.Context) (string) {
    status := API("/api/web/domain/delete", ctx)
    
    //vhost_id := util.Query(ctx, "vhost_id")
    
    
    if status == "success" {
        set_error("Deleted domain successfully!", ctx)
        ctx.Redirect("/websites", 302)
        return "did it!"
    }
    
    set_error("Failed to delete domain. Error given: " + status, ctx)
    ctx.Redirect("/websites", 302)
    
    return "Failed to add domain. Error given: " + status
    
}

func addsite(ctx *macaron.Context) (string){
    status := API("/api/web/domain/add", ctx)
    
    domainname := util.Query(ctx, "domainname")
    
    if status == "success" {
        set_error("Added " + domainname + " successfully!", ctx)
        ctx.Redirect("/websites", 302)
        return "did it!"
    }
    
    set_error("Failed to add domain. Error given: " + status, ctx)
    ctx.Redirect("/websites", 302)
    
    return "Failed to add domain. Error given: " + status
}

func websites(ctx *macaron.Context) (string){
	_, auth := util.Auth(ctx, "websites")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/websites.tpl")
	tpl.Parse("websites")
	
	websites := API("/api/web/domain/list", ctx)
	
	domains := make(map[string]map[string]string)
	json.Unmarshal([]byte(websites), &domains)
	
	for _, domain := range domains {
        tpl.Assign("vhost_id", domain["vhost_id"])
        tpl.Assign("system_username", domain["system_username"])
        tpl.Assign("domain", domain["domain"])
        tpl.Assign("documentroot", domain["documentroot"])
        tpl.Assign("ipaddr", domain["ipaddr"])
        tpl.Assign("ssl_enabled", domain["ssl_enabled"])
        
        determin_fm_page := strings.Split(domain["documentroot"], "www")
        filemanager_path := "www" + determin_fm_page[1]
        
        tpl.Assign("filemanager_path", filemanager_path)
	    tpl.Parse("websites/domain")
	}
	
	
	return header(ctx) + tpl.Out() + footer(ctx)
}


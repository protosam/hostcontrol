package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
)


func init(){
	route("/dns", dns, "BOTH")
	route("/dns/domain/add", adddomain, "BOTH")
	route("/dns/domain/delete", deletedomain, "BOTH")
	route("/dns/record/add", addrecord, "BOTH")
	route("/dns/record/delete", deleterecord, "BOTH")
	route("/dns/record/edit", editrecord, "BOTH")
}


func editrecord(ctx *macaron.Context) (string){
    status := API("/api/dns/record/edit", ctx)
    
    
    if status == "success" {
        set_error("Edited record successfully!", ctx)
        ctx.Redirect("/dns", 302)
        return "did it!"
    }
    
    set_error("Failed to edit record. Error given: " + status, ctx)
    ctx.Redirect("/dns", 302)
    
    return "Failed to edit record. Error given: " + status
}

func addrecord(ctx *macaron.Context) (string){
    status := API("/api/dns/record/add", ctx)
    
    
    if status == "success" {
        set_error("Added record successfully!", ctx)
        ctx.Redirect("/dns", 302)
        return "did it!"
    }
    
    set_error("Failed to add record. Error given: " + status, ctx)
    ctx.Redirect("/dns", 302)
    
    return "Failed to add record. Error given: " + status
}

func deleterecord(ctx *macaron.Context) (string){
    status := API("/api/dns/record/delete", ctx)
    
    
    if status == "success" {
        set_error("Deleted record successfully!", ctx)
        ctx.Redirect("/dns", 302)
        return "did it!"
    }
    
    set_error("Failed to delete record. Error given: " + status, ctx)
    ctx.Redirect("/dns", 302)
    
    return "Failed to delete record. Error given: " + status
}

func deletedomain(ctx *macaron.Context) (string){
    status := API("/api/dns/domain/delete", ctx)
    
    domainname := util.Query(ctx, "domain")
    
    if status == "success" {
        set_error("Deleted " + domainname + " successfully!", ctx)
        ctx.Redirect("/dns", 302)
        return "did it!"
    }
    
    set_error("Failed to delete domain. Error given: " + status, ctx)
    ctx.Redirect("/dns", 302)
    
    return "Failed to delete domain. Error given: " + status
}

func adddomain(ctx *macaron.Context) (string){
    status := API("/api/dns/domain/add", ctx)
    
    domainname := util.Query(ctx, "domain")
    
    if status == "success" {
        set_error("Added " + domainname + " successfully!", ctx)
        ctx.Redirect("/dns", 302)
        return "did it!"
    }
    
    set_error("Failed to add domain. Error given: " + status, ctx)
    ctx.Redirect("/websites", 302)
    
    return "Failed to add domain. Error given: " + status
}

func dns(ctx *macaron.Context) (string){
	_, auth := util.Auth(ctx, "dns")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/dns.tpl")
	tpl.Parse("dns")
	


	// list domains and records
	dns_data := API("/api/dns/list", ctx)

	
	// map[domain]map[record_id]map[key]value
	data := make(map[string]map[string]map[string]string)
	json.Unmarshal([]byte(dns_data), &data)
	
	for domain, records := range data {
        tpl.Assign("domain_name", domain)
	    tpl.Parse("dns/domain")
	    
	    for key, record := range records {
	        if key == "placebo" {
	            continue
	        }
            tpl.Assign("record_change_date", record["record_change_date"])
            tpl.Assign("record_content", record["record_content"])
            tpl.Assign("record_disabled", record["record_disabled"])
            tpl.Assign("record_domain_id", record["record_domain_id"])
            tpl.Assign("record_id", record["record_id"])
            tpl.Assign("record_name", record["record_name"])
            tpl.Assign("record_ordername", record["record_ordername"])
            tpl.Assign("record_prio", record["record_prio"])
            tpl.Assign("record_ttl", record["record_ttl"])
            tpl.Assign("record_type", record["record_type"])
    	    tpl.Parse("dns/domain/record")
	    }
	}

	return header(ctx) + tpl.Out() + footer(ctx)
}


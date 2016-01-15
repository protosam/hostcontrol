package rhel7

import (
//	"fmt"
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"

    "encoding/json"
	"time"
	"strconv"
)


// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func Dns(ctx *macaron.Context) (string) {
	return "mydns"
}

func DnsAddDomain(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "dns")
	
	if ! auth {
		return "not_authorized"
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	
	domain := util.Query(ctx, "domain")
	
	if domain == "" {
	    return "domain_required"
	}
	
    db, err := util.MySQL();
    defer db.Close()
    
	xstmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`domains` set `id`=NULL, `name`=?, `master`=NULL, `last_check`=NULL, `type`='NATIVE', `notified_serial`=?, `account`=?")
	
	res, err := xstmt.Exec(domain, timestamp, hcuser.System_username)
	xstmt.Close()

	if err != nil {
	    return "failed_to_create_domain"
	}
    
    inserted_id, err := res.LastInsertId()
	if err != nil {
	    return "failed_to_create_domain"
	}
    
	ystmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`records` set `id`=NULL, `domain_id`=?, `name`=?, `type`='SOA', `content`=?, `ttl`='86400', `prio`='0', `change_date`=?, `disabled`='0', `ordername`='0', `auth`='1'")
	_, yerr := ystmt.Exec(inserted_id, domain, "localhost webmaster@localhost 1", timestamp)
	ystmt.Close()
	if yerr != nil {
	    return "failed_to_create_soa"
	}
	
	return "success"
}

func DnsDeleteDomain(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "dns")
	
	if ! auth {
		return "not_authorized"
	}
	
	domain := util.Query(ctx, "domain")
	
	if domain == "" {
	    return "domain_required"
	}
	
    db, _ := util.MySQL();
    defer db.Close()
    
	xstmt, _ := db.Prepare("DELETE FROM `hostcontrol`.`domains` where `name`=? and `account`=?")
	
	_, err := xstmt.Exec(domain, hcuser.System_username)
	xstmt.Close()

	if err != nil {
	    return "failed_to_delete_domain"
	}
	
	return "success"
}

func DnsAddRecord(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "dns")
	
	if ! auth {
		return "not_authorized"
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	
	// http://panel.just.ninja/api/dns/record/add?domain=helloworld.com&name=helloworld.com&content=1.2.3.4&type=a&ttl=300&priority=0
	
	domain := util.Query(ctx, "domain")
	if domain == "" {
	    return "domain_required"
	}
	name := util.Query(ctx, "name")
	if name == "" {
	    return "name_required"
	}
	content := util.Query(ctx, "content")
	if content == "" {
	    return "content_required"
	}
	rtype := util.Query(ctx, "type")
	if rtype == "" {
	    return "type_required"
	}
	ttl := util.Query(ctx, "ttl")
	if ttl == "" {
	    return "ttl_required"
	}
	priority := util.Query(ctx, "priority")
	if priority == "" {
	    return "priority_required"
	}
	
	
	
    db, err := util.MySQL();
    defer db.Close()
	
    // lookup domain
    stmt, _ := db.Prepare("select * from `hostcontrol`.`domains` where `name`=? and `account`=?")
    row, err := stmt.Query(domain, hcuser.System_username)
    if err != nil {
        return "domain_not_found"
    }
    var domain_id string
    var domain_name string
    var domain_master string
    var domain_last_check string
    var domain_type string
    var domain_notified_serial string
    var domain_account string
    
    
    if row.Next() {
        row.Scan(&domain_id, &domain_name, &domain_master, &domain_last_check, &domain_type, &domain_notified_serial, &domain_account)
    }else{
        return "failed_to_find_domain"
    }
    
	ystmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`records` set `id`=NULL, `domain_id`=?, `name`=?, `type`=?, `content`=?, `ttl`=?, `prio`=?, `change_date`=?, `disabled`='0', `ordername`='0', `auth`='1'")
	_, yerr := ystmt.Exec(domain_id, name, rtype, content, ttl, priority, timestamp)
	ystmt.Close()
	if yerr != nil {
	    return "failed_to_create_record"
	}
	
	// update serial for domain
	ustmt, _ := db.Prepare("UPDATE `hostcontrol`.`domains` SET `notified_serial`=? WHERE `name`=? and `account`=?")
	ustmt.Exec(timestamp, name, hcuser)
	ustmt.Close()
	
	return "success"
}


func DnsEditRecord(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "dns")
	
	if ! auth {
		return "not_authorized"
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	
	// http://panel.just.ninja/api/dns/record/add?domain=helloworld.com&name=helloworld.com&content=1.2.3.4&type=a&ttl=300&priority=0
	

	record_id := util.Query(ctx, "record_id")
	if record_id == "" {
	    return "record_id_required"
	}
	name := util.Query(ctx, "name")
	if name == "" {
	    return "name_required"
	}
	content := util.Query(ctx, "content")
	if content == "" {
	    return "content_required"
	}
	rtype := util.Query(ctx, "type")
	if rtype == "" {
	    return "type_required"
	}
	ttl := util.Query(ctx, "ttl")
	if ttl == "" {
	    return "ttl_required"
	}
	priority := util.Query(ctx, "priority")
	if priority == "" {
	    return "priority_required"
	}
	
	
    db, _ := util.MySQL();
    defer db.Close()
    
    
    // confirm user owns record
    stmt, _ := db.Prepare("SELECT * FROM records LEFT JOIN domains ON records.domain_id=domains.id WHERE `records`.`id`=? and `account`=?")
    rows, _ := stmt.Query(record_id, hcuser.System_username)
    stmt.Close()

    if ! rows.Next() {
        return "record_lookup_failed"
    }
    
    
    
	xstmt, err := db.Prepare("update records set name=?, type=?, content=?, ttl=?, prio=?, change_date=? where id=?")
    if err != nil {
	    return "failed_to_update_record"
	}
	
	_, xerr := xstmt.Exec(name, rtype, content, ttl, priority, timestamp, record_id)
	xstmt.Close()

	if xerr != nil {
	    return "failed_to_update_record"
	}
	
	return "success"
}

func DnsDeleteRecord(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "dns")
	
	if ! auth {
		return "not_authorized"
	}
	
	record_id := util.Query(ctx, "record_id")
	
	if record_id == "" {
	    return "record_id_required"
	}
	
    db, _ := util.MySQL();
    defer db.Close()
    
    
   // confirm user owns record
    stmt, _ := db.Prepare("SELECT * FROM records LEFT JOIN domains ON records.domain_id=domains.id WHERE `records`.`id`=? and `account`=?")
    rows, _ := stmt.Query(record_id, hcuser.System_username)
    stmt.Close()

    if ! rows.Next() {
        return "record_lookup_failed"
    }
    
	xstmt, err := db.Prepare("DELETE FROM `hostcontrol`.`records` where `id`=?")
	if err != nil {
	    return "failed_to_delete_record" + string(err.Error())
	}
	_, xerr := xstmt.Exec(record_id)
	xstmt.Close()

	if xerr != nil {
	    return "failed_to_delete_record"
	}
	
	return "success"
}

func DnsList(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "dns")
	
	if ! auth {
		return "not_authorized"
	}


    db, _ := util.MySQL();
    defer db.Close()
    // alternative by domain: SELECT * FROM domains RIGHT JOIN records ON domains.id=records.domain_id WHERE domains.name=?
    stmt, _ := db.Prepare("SELECT domains.id,domains.name, domains.type, domains.notified_serial, domains.account, records.id, records.domain_id, records.name, records.type, records.content, records.ttl, records.prio, records.change_date, records.disabled, records.ordername, records.auth FROM domains LEFT JOIN records ON records.domain_id=domains.id WHERE `account`=?")
    rows, _ := stmt.Query(hcuser.System_username)
    stmt.Close()
    
    
    // map[domain]map[record_id]map[key]value
    data := make(map[string]map[string]map[string]string)
    for rows.Next() {
        var domain_id string
        var domain_name string
        var domain_type string
        var domain_notified_serial string
        var domain_account string
        var record_id string
        var record_domain_id string
        var record_name string
        var record_type string
        var record_content string
        var record_ttl string
        var record_prio string
        var record_change_date string
        var record_disabled string
        var record_ordername string
        var record_auth string
        
        
        // if any of these have a NULL returned... we are fucked. -_-
        rows.Scan(&domain_id, &domain_name, &domain_type, &domain_notified_serial, &domain_account, &record_id, &record_domain_id, &record_name, &record_type, &record_content, &record_ttl, &record_prio, &record_change_date, &record_disabled, &record_ordername, &record_auth)
        
        if _, ok := data[domain_name]; ! ok {
            data[domain_name] = make(map[string]map[string]string)
        }
        
        if record_name != "" {
            data[domain_name][record_id] = make(map[string]string)
            data[domain_name][record_id]["record_id"] = record_id
            data[domain_name][record_id]["record_domain_id"] = record_domain_id
            data[domain_name][record_id]["record_name"] = record_name
            data[domain_name][record_id]["record_type"] = record_type
            data[domain_name][record_id]["record_content"] = record_content
            data[domain_name][record_id]["record_ttl"] = record_ttl
            data[domain_name][record_id]["record_prio"] = record_prio
            data[domain_name][record_id]["record_change_date"] = record_change_date
            data[domain_name][record_id]["record_disabled"] = record_disabled
            data[domain_name][record_id]["record_ordername"] = record_ordername
        }
            
        data[domain_name]["placebo"] = make(map[string]string)
        data[domain_name]["placebo"]["placebo"] = "placebo"
    
    }
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}
	return string(output)
}
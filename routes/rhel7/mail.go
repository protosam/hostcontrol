package rhel7

import (
//	"fmt"
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
	"strings"
	"os"
	"encoding/json"
)



// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func Mail(ctx *macaron.Context) (string) {
	return"dovecot/postfix"
}

func MailList(ctx *macaron.Context) (string) {
    hcuser, auth := util.Auth(ctx, "mail")
	
	if ! auth {
		return "not_authorized"
	}
	
	
    db, _ := util.MySQL();
    defer db.Close()

    stmt, _ := db.Prepare("SELECT mail_domains.domain, email_id, email FROM mail_domains LEFT JOIN mail_users ON mail_users.domain=mail_domains.domain WHERE `system_username`=?")
    rows, _ := stmt.Query(hcuser.System_username)
    stmt.Close()
    
    // map[domain]map[email_id]map[key]value
    data := make(map[string]map[string]map[string]string)
    for rows.Next() {
        var email_id string
        var email string
        var domain string
        
        rows.Scan(&domain, &email_id, &email)
        if _, ok := data[domain]; ! ok {
            data[domain] = make(map[string]map[string]string)
        }
        
        
        // some logic for if it's empty or not...
        if email_id != "" && email != "" {
            data[domain][email_id] = make(map[string]string)
            data[domain][email_id]["email_id"] = email_id
            data[domain][email_id]["email"] = email
        }
        
        data[domain]["placebo"] = make(map[string]string)
        data[domain]["placebo"]["placebo"] = "placebo"
        
    }
    
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}
	return string(output)
}

func MailAddDomain(ctx *macaron.Context) (string) {
    hcuser, auth := util.Auth(ctx, "mail")
	
	if ! auth {
		return "not_authorized"
	}
	
	domain := util.Query(ctx, "domain")
	
	if domain == "" {
	    return "domain_required"
	}
	
    db, _ := util.MySQL();
    defer db.Close()
    
	xstmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`mail_domains` set `domain_id`=NULL, `domain`=?, `system_username`=?")
	
	_, err := xstmt.Exec(domain, hcuser.System_username)
	xstmt.Close()

	if err != nil {
	    return "failed_to_create_domain"
	}

    return "success"
}



func MailDeleteDomain(ctx *macaron.Context) (string) {
    hcuser, auth := util.Auth(ctx, "mail")
	
	if ! auth {
		return "not_authorized"
	}
	
	domain := util.Query(ctx, "domain")
	
	if domain == "" {
	    return "domain_required"
	}
	
    db, _ := util.MySQL();
    defer db.Close()
    
	xstmt, _ := db.Prepare("DELETE FROM `hostcontrol`.`mail_domains` WHERE `domain`=? AND `system_username`=?")
	
	_, err := xstmt.Exec(domain, hcuser.System_username)
	xstmt.Close()

	if err != nil {
	    return "failed_to_delete_domain"
	}
	
	os.RemoveAll("/home/vmail/" + domain)

    return "success"
}




func MailAddUser(ctx *macaron.Context) (string) {
    hcuser, auth := util.Auth(ctx, "mail")
	
	if ! auth {
		return "not_authorized"
	}
	
	domain := util.Query(ctx, "domain")
	if domain == "" {
	    return "domain_required"
	}
	username := util.Query(ctx, "username")
	if username == "" {
	    return "username_required"
	}
	password := util.Query(ctx, "password")
	if password == "" {
	    return "password_required"
	}
	
	email_address := username + "@" + domain
	
    db, _ := util.MySQL();
    defer db.Close()
    
    // check if user owns domain
    dstmt, _ := db.Prepare("SELECT * FROM `hostcontrol`.`mail_domains` WHERE `domain`=? and `system_username`=?")
    row1, _ := dstmt.Query(domain, hcuser.System_username)
    defer dstmt.Close()
    if ! row1.Next(){
        return "domain_not_found"
    }
    
    // make sure email address does not already exist
    estmt, _ := db.Prepare("SELECT * FROM `hostcontrol`.`mail_users` WHERE email=? and domain=?")
    row2, _ := estmt.Query(email_address, domain)
    defer estmt.Close()
    if row2.Next(){
        return "email_account_exists"
    }
    
	xstmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`mail_users` set `email`=?, `password`=ENCRYPT(?), `domain`=?")
	_, err := xstmt.Exec(email_address, password, domain)
	xstmt.Close()

	if err != nil {
	    return "failed_to_create_domain"
	}

    return "success"
}


func MailEditUser(ctx *macaron.Context) (string) {
    hcuser, auth := util.Auth(ctx, "mail")
	
	if ! auth {
		return "not_authorized"
	}
	
	email_address := util.Query(ctx, "email")
	if email_address == "" {
	    return "email_required"
	}
	password := util.Query(ctx, "password")
	if email_address == "" {
	    return "password_required"
	}
	
	strsplt := strings.Split(email_address, "@")
    if len(strsplt) != 2 {
        return "invalid_email"
    }
    
    //username := strsplt[0]
    domain := strsplt[1]
    
    db, _ := util.MySQL();
    defer db.Close()
    
    // check if user owns domain
    dstmt, _ := db.Prepare("SELECT * FROM `hostcontrol`.`mail_domains` WHERE `domain`=? and `system_username`=?")
    row1, _ := dstmt.Query(domain, hcuser.System_username)
    defer dstmt.Close()
    if ! row1.Next(){
        return "domain_not_found"
    }
    
    
	// update serial for domain
	ustmt, _ := db.Prepare("UPDATE `hostcontrol`.`mail_users` SET `password`=ENCRYPT(?) WHERE `email`=?")
	ustmt.Exec(password, email_address)
	ustmt.Close()
    
    return "success"
}



func MailUserDelete(ctx *macaron.Context) (string) {
    hcuser, auth := util.Auth(ctx, "mail")
	
	if ! auth {
		return "not_authorized"
	}
	
	email_address := util.Query(ctx, "email")
	if email_address == "" {
	    return "email_required"
	}
	
	strsplt := strings.Split(email_address, "@")
    if len(strsplt) != 2 {
        return "invalid_email"
    }
    
    username := strsplt[0]
    domain := strsplt[1]
    
    db, _ := util.MySQL();
    defer db.Close()
    
    // check if user owns domain
    dstmt, _ := db.Prepare("SELECT * FROM `hostcontrol`.`mail_domains` WHERE `domain`=? and `system_username`=?")
    row1, _ := dstmt.Query(domain, hcuser.System_username)
    defer dstmt.Close()
    if ! row1.Next(){
        return "domain_not_found"
    }
    
    os.RemoveAll("/home/vmail/" + domain + "/" + username)
    
	// update serial for domain
	ustmt, _ := db.Prepare("DELETE FROM `hostcontrol`.`mail_users` WHERE `email`=?")
	ustmt.Exec(email_address)
	ustmt.Close()
    
    return "success"
}

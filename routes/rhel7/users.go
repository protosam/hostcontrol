package rhel7

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
	"os/exec"
	"os/user"
	
	"encoding/json"
	"strings"
)




// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func Distro(ctx *macaron.Context) (string) {
	return "RHEL7"
}


func Adduser(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "sysusers")
	if ! auth {
		return "not_authorized"
	}

	username := util.Query(ctx, "username")
	password := util.Query(ctx, "password")

    if username == "" || username == "root" {
        return "username_required"
    }

    if password == "" {
        return "password_required"
    }

	db, _ := util.MySQL()
	defer db.Close()
	
	// check if username is available
    _, lookup_err1 := user.Lookup(username)
    if lookup_err1 == nil {
        return "username_taken"
    }
	
	// add the user
	util.Cmd("useradd", []string{username,"-d", "/home/"+username})

	// make sure user was added
    _, lookup_err2 := user.Lookup(username)
    if lookup_err2 != nil {
        return "unable_to_create"
    }
    
    // set the password
	util.Bash("echo " + util.SHSanitize(password) + " | passwd " + util.SHSanitize(username) + " --stdin")
    
	new_token := util.MkToken()
	
	// add the user
	istmt, _ := db.Prepare("insert hostcontrol_users set hostcontrol_id=null, system_username=?, privileges=?, owned_by=?, login_token=?, email_address=?")

    privileges := ""
    
    perm_all := util.Query(ctx, "allperms")
    if strings.Contains(hcuser.Privileges, "all") && perm_all != "" {
    	privileges += "all "
	}
    perm_websites := util.Query(ctx, "websites")
	if (strings.Contains(hcuser.Privileges, "websites") || strings.Contains(hcuser.Privileges, "all")) && perm_websites != "" {
	    privileges += "websites "
	}
    perm_mail := util.Query(ctx, "mail")
	if (strings.Contains(hcuser.Privileges, "mail") || strings.Contains(hcuser.Privileges, "all")) && perm_mail != "" {
	    privileges += "mail "
	}
    perm_databases := util.Query(ctx, "databases")
	if (strings.Contains(hcuser.Privileges, "databases") || strings.Contains(hcuser.Privileges, "all")) && perm_databases != "" {
	    privileges += "databases "
	}
    perm_ftpusers := util.Query(ctx, "ftpusers")
	if (strings.Contains(hcuser.Privileges, "ftpusers") || strings.Contains(hcuser.Privileges, "all")) && perm_ftpusers != "" {
	    privileges += "ftpusers "
	}
    perm_dns := util.Query(ctx, "dns")
	if (strings.Contains(hcuser.Privileges, "dns") || strings.Contains(hcuser.Privileges, "all")) && perm_dns != "" {
	    privileges += "dns "
	}
    perm_sysusers := util.Query(ctx, "sysusers")
	if (strings.Contains(hcuser.Privileges, "sysusers") || strings.Contains(hcuser.Privileges, "all")) && perm_sysusers != "" {
	    privileges += "sysusers "
	}

	istmt.Exec(username,privileges,hcuser.System_username,new_token,"")
	istmt.Close()


	return "success"
}


func cleanupuserdata(username string, ctx *macaron.Context){
        util.Sudo(username, ctx)
    
        // websites
        websites := make(map[string]map[string]string)
        webdata := API("/api/web/domain/list", ctx)
        json.Unmarshal([]byte(webdata), &websites)
        for _, info := range websites {
            API("/api/web/domain/delete?vhost_id="+info["vhost_id"], ctx)
        }

        // dns
        dnsdomains := make(map[string]map[string]map[string]string)
        dnsdata := API("/api/dns/list", ctx)
        json.Unmarshal([]byte(dnsdata), &dnsdomains)
        for domain, _ := range dnsdomains {
            API("/api/dns/domain/delete?domain=" + domain, ctx)
        }
        
        // ftpusers
        ftpusers := make(map[string]map[string]string)
        ftpdata := API("/api/ftpusers/list", ctx)
        json.Unmarshal([]byte(ftpdata), &ftpusers)
        for _, ftpuser := range ftpusers {
            API("/api/ftpusers/delete?ftpuser=" + ftpuser["username"], ctx)
        }

        // mail
        maildomains := make(map[string]map[string]map[string]string)
        maildata := API("/api/mail/list", ctx)
        json.Unmarshal([]byte(maildata), &maildomains)
        for maildomain, _ := range maildomains {
            if maildomain != "placebo" {
                API("/api/mail/domain/delete?domain=" + maildomain, ctx)
            }
        }


        // databases and grants
        var databases []string
        dbdata := API("/api/sql/databases/list", ctx)
        json.Unmarshal([]byte(dbdata), &databases)
        for _, db_name := range databases {
	        // list database grants
	        var grants []string
	        grantdata := API("/api/sql/grants/list", ctx)
	        json.Unmarshal([]byte(grantdata), &grants)
	        for _, db_user := range grants {
	            API("/api/sql/grants/delete?db_name=" + db_name + "&db_user=" + db_user, ctx)
	        }
            API("/api/sql/databases/delete?db_name=" + db_name, ctx)
        }
        
        // database users
        var dbusers []string
        dbuserdata := API("/api/sql/users/list", ctx)
        json.Unmarshal([]byte(dbuserdata), &dbusers)
        for _, username := range dbusers {
            API("/api/sql/users/delete?db_user=" + username, ctx)
        }

    util.Sudo("", ctx)
}

func Deleteuser(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "sysusers")
	if ! auth {
		return "not_authorized"
	}

	username := util.Query(ctx, "username")
	
    if username == "" || username == "root" {
        return "username_required"
    }


	db, _ := util.MySQL()
	defer db.Close()
	
	// check if user actually owns child
	if ! util.ChkPaternity(hcuser.System_username, username) {
        return "failed_ownership_check"
	}
	
	
	
    users := make(map[string]map[string]string)
    users = util.Getusers(username, users, db)
    for _, subuser := range users {
        cleanupuserdata(subuser["system_username"], ctx)
    	// delete the user and homedir
    	util.Cmd("userdel", []string{subuser["system_username"],"-f", "-r"})
    	// remove the user
        stmt, _ := db.Prepare("delete from hostcontrol_users where system_username=?")
        stmt.Exec(subuser["system_username"])
        stmt.Close()
        
    }
    
    cleanupuserdata(username, ctx)
    
	// delete the user and homedir
	util.Cmd("userdel", []string{username,"-f", "-r"})

	// make sure user was delete
    _, lookup_err2 := user.Lookup(username)
    if lookup_err2 == nil {
        return "failed_to_delete_user"
    }
	
	// remove the user
    stmt, _ := db.Prepare("delete from hostcontrol_users where system_username=?")
    stmt.Exec(username)
    stmt.Close()


	return "success"
}

func Listusers(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "sysusers")
	
	if ! auth {
		return "not_authorized"
	}
	
    db, _ := util.MySQL();
	defer db.Close()
	
	
    data := make(map[string]map[string]string)
    data = util.Getusers(hcuser.System_username, data, db)


    
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}

	return string(output)
}



func ChPassword(username string, password string) (bool) {
    if username == "" {
        return false
    }

    if password == "" {
        return false
    }
    // set the password
	util.Bash("echo " + util.SHSanitize(password) + " | passwd " + util.SHSanitize(username) + " --stdin")

    return true
}

func ChkLogin(username string, password string) bool {
	args := []string{"-u", username, "-p", password, "-s", "password-auth"}
	_, err := exec.Command("testsaslauthd", args...).Output();
	if err != nil {
		return false
	}

	return true
}

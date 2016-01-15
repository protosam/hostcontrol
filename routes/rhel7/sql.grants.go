package rhel7

import (
//    "fmt"
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
	"strings"
)



// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func SqlGrantsList(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}


    db_name := util.Query(ctx, "db_name")

    if db_name == "" {
        return "db_name_required"
    }
    
    owner := strings.Split(db_name, "_")[0]
    
    if owner != hcuser.System_username {
        return "failed_not_yours"
    }
	
	
    db, _ := util.MySQL();
    defer db.Close()
    
    stmt, _ := db.Prepare("select user from mysql.db where db=?")
    rows, err := stmt.Query(db_name)
    if err != nil {
        return "bad_characters_used "
    }    
    stmt.Close()
    var data []string
    
    for rows.Next() {
        var db_user string
        
        rows.Scan(&db_user)
        
        data = append(data, db_user)
    }
    
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}


	return string(output)
}

func SqlGrantsAdd(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	db_name := util.Query(ctx, "db_name")
	

    if db_name == "" {
        return "db_name_required"
    }
    
	username := util.Query(ctx, "db_user")
	

    if username == "" {
        return "username_required"
    }
    
    dbowner := strings.Split(db_name, "_")[0]
    userowner := strings.Split(username, "_")[0]
    
    if dbowner != hcuser.System_username || userowner != hcuser.System_username {
        return "failed_not_yours"
    }
    

	
    db, _ := util.MySQL();
    defer db.Close()
    
    db_name = util.LastResortSanitize(db_name)
    username = util.LastResortSanitize(username)

    _, err := db.Exec("GRANT ALL ON "+db_name+".* TO '"+username+"'@'%';")
    if err != nil {
        
        return "failed_to_create_grant"
    }
	return "success"
}

func SqlGrantsDelete(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
    db_name := util.Query(ctx, "db_name")

    if db_name == "" {
        return "db_name_required"
    }
    
	username := util.Query(ctx, "db_user")
	

    if username == "" {
        return "username_required"
    }
    
    dbowner := strings.Split(db_name, "_")[0]
    userowner := strings.Split(username, "_")[0]
    
    if dbowner != hcuser.System_username || userowner != hcuser.System_username {
        return "failed_not_yours"
    }
    

	
    db, _ := util.MySQL();
    defer db.Close()
    
    db_name = util.LastResortSanitize(db_name)
    username = util.LastResortSanitize(username)

    _, err := db.Exec("REVOKE ALL ON "+db_name+".* FROM '"+username+"'@'%';")
    if err != nil {
        
        return "failed_to_delete_grant"
    }
	return "success"
}

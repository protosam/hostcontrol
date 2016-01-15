package rhel7

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
	"strings"
)



// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func SqlDatabasesList(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	
    db, _ := util.MySQL();
    defer db.Close()
    
    rows, err := db.Query("show databases like '" + hcuser.System_username + "\\_%'")
    if err != nil {
        return "bad_characters_used "
    }    
    
    var data []string
    
    for rows.Next() {
        var db_name string
        
        rows.Scan(&db_name)
        
        data = append(data, db_name)
    }
    
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}

	return string(output)
}

func SqlDatabasesAdd(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	db_name := util.Query(ctx, "db_name")
	

    if db_name == "" {
        return "db_name_required"
    }
    

	
    db, _ := util.MySQL();
    defer db.Close()
    
//    stmt, _ := db.Prepare("CREATE USER ?@'%' IDENTIFIED BY ?;")
//    _, err := stmt.Exec(hcuser.System_username + "_" + username, password)
    db_name = util.LastResortSanitize(db_name)
    db_name = string(hcuser.System_username + "_" + db_name)

    stmt, err := db.Prepare("create database "+db_name+"")
    if err != nil {
        return "bad_characters_used "
    }
    _, err = stmt.Exec()
    if err != nil {
        return "failed_to_create_database"
    }
    stmt.Close()
    
	return "success"
}

func SqlDatabasesDelete(ctx *macaron.Context) (string) {
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
    
//    stmt, _ := db.Prepare("CREATE USER ?@'%' IDENTIFIED BY ?;")
//    _, err := stmt.Exec(hcuser.System_username + "_" + username, password)
    db_name = util.LastResortSanitize(db_name)

    stmt, err := db.Prepare("drop database "+db_name+"")
    if err != nil {
        return "bad_characters_used "
    }
    _, err = stmt.Exec()
    if err != nil {
        return "failed_to_create_database"
    }
    stmt.Close()
    
	return "success"
}

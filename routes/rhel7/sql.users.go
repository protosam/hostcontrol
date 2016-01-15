package rhel7

import (
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
	"encoding/json"
	"strings"
)

func SqlUsersList(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	
    db, _ := util.MySQL();
    defer db.Close()
    
    stmt, _ := db.Prepare("select DISTINCT user from mysql.user where user like concat(?,'_%')")
    rows, err := stmt.Query(hcuser.System_username)
    if err != nil {
        return "failed_user_select_query" + string(err.Error())
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


func SqlUsersAdd(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	username := util.Query(ctx, "db_user")
	password := util.Query(ctx, "password")
	

    if username == "" {
        return "username_required"
    }
    
    if password == "" {
        return "password_required"
    }

	
    db, _ := util.MySQL();
    defer db.Close()
    
//    stmt, _ := db.Prepare("CREATE USER ?@'%' IDENTIFIED BY ?;")
//    _, err := stmt.Exec(hcuser.System_username + "_" + username, password)
    db_user := string(hcuser.System_username + "_" + username)
    
    db_user = util.LastResortSanitize(db_user)
    password = util.LastResortSanitize(password)

    stmt, err := db.Prepare("CREATE USER '"+db_user+"'@'%' IDENTIFIED BY '"+password+"'")
    if err != nil {
        return "bad_characters_used "
    }
    _, err = stmt.Exec()
    if err != nil {
        return "failed_to_create_user"
    }
    stmt.Close()
    
	return "success"
}




func SqlUsersDelete(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	username := util.Query(ctx, "db_user")
	

    if username == "" {
        return "username_required"
    }
    
    owner := strings.Split(username, "_")[0]
    
    if owner != hcuser.System_username {
        return "failed_not_yours"
    }
	
    db, _ := util.MySQL();
    defer db.Close()
    
    
    db_user := util.LastResortSanitize(username)
    //password = strings.Replace(password, "'", "\\'", -1)

    stmt, err := db.Prepare("DROP USER '"+db_user+"'")
    if err != nil {
        return "bad_characters_used"
    }
    _, err = stmt.Exec()
    if err != nil {
        return "failed_to_delete_user"
    }
    stmt.Close()
    
	return "success"
}





func SqlUsersEdit(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "databases")
	
	if ! auth {
		return "not_authorized"
	}
	
	username := util.Query(ctx, "db_user")
	password := util.Query(ctx, "password")
    owner := strings.Split(username, "_")[0]
	

    if username == "" {
        return "db_user_required"
    }
    
    if password == "" {
        return "password_required"
    }

    
    if owner != hcuser.System_username {
        return "failed_not_yours"
    }
	
    db, _ := util.MySQL();
    defer db.Close()
    
    db_user := util.LastResortSanitize(username)
    password = util.LastResortSanitize(password)

    
    _, err := db.Exec("SET PASSWORD FOR '"+db_user+"' = PASSWORD('"+password+"');")
    if err != nil {
        return "bad_characters_used "
    }
    
	return "success"
}

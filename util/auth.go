package util

import (
	"gopkg.in/macaron.v1"
	"os"
	"syscall"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"os/user"
	"strings"
)

type API_function func(ctx *macaron.Context) (string)

type User struct {
	Hostcontrol_id int
	System_username string
	Privileges string
	Owned_by string
	Login_token string
	Email_address string
	HomeDir string
	Sudo bool
}

func Auth(ctx *macaron.Context, permission string) (User,bool) {
	//api_authid := string(ctx.Query("api_authid"))
	//api_token := string(ctx.Query("api_token"))

	ckhostcontrol_id := ctx.GetCookie("hostcontrol_id")
	cklogin_token := ctx.GetCookie("login_token")
	
	api_token_user_id   := ""
	api_token_id        := ""
	
	raw_api_token := Query(ctx, "api_token")
	if raw_api_token != "" {
	    parsed_token := strings.Split(raw_api_token, "/")
	    if len(parsed_token) == 2 {
	        api_token_user_id = parsed_token[0]
	        api_token_id = parsed_token[1]
	    }
	}

	var hcuser User
	hcuser.Sudo = false

	db, _ := MySQL()
	defer db.Close()
	
    if api_token_user_id != "" {
    	tstmt, err := db.Prepare("SELECT * from hostcontrol_user_tokens LEFT JOIN hostcontrol_users ON hostcontrol_user_tokens.hostcontrol_id=hostcontrol_users.hostcontrol_id WHERE hostcontrol_user_tokens.hostcontrol_id=? and token=?")
    	if err != nil {
    		fmt.Println(string(err.Error()))
    	}
    	trows, _ := tstmt.Query(api_token_user_id, api_token_id)
    	tstmt.Close()
    
    	// check if we have a row returned...
    	if trows.Next() {
    	    var token string
    	    var hostcontrol_id string
    	    var description string
    	    var token_id string
    		trows.Scan(&token, &hostcontrol_id, &description, &token_id, &hcuser.Hostcontrol_id, &hcuser.System_username, &hcuser.Privileges, &hcuser.Owned_by, &hcuser.Login_token, &hcuser.Email_address)


            if strings.Contains(hcuser.Privileges, permission) || permission == "any" || strings.Contains(hcuser.Privileges, "all")  {
		        return hcuser, true
            }
	    }
    }
    
    
    // token authentication check
	stmt, err := db.Prepare("SELECT * from hostcontrol_users WHERE hostcontrol_id = ? and login_token = ?")
	if err != nil {
		fmt.Println(string(err.Error()))
	}
	rows, _ := stmt.Query(ckhostcontrol_id, cklogin_token)
	stmt.Close()

	// check if we have a row returned...
	if rows.Next() {
		rows.Scan(&hcuser.Hostcontrol_id, &hcuser.System_username, &hcuser.Privileges, &hcuser.Owned_by, &hcuser.Login_token, &hcuser.Email_address)
	}else{
		return hcuser, false
	}

    // check for sudo cookie
	cksudo := ctx.GetCookie("sudo")
	
	bkmount := string(ctx.Params("sudo"))
	if bkmount != "" {
	    cksudo = bkmount
	}
	
	if (cksudo != "" && ChkPaternity(hcuser.System_username, cksudo)) || bkmount != "" {
    	stmt, _ = db.Prepare("SELECT * from hostcontrol_users WHERE system_username = ?")
    	rows, _ = stmt.Query(cksudo)
    	stmt.Close()
    
    	// check if we have a row returned...
    	if rows.Next() {
    		rows.Scan(&hcuser.Hostcontrol_id, &hcuser.System_username, &hcuser.Privileges, &hcuser.Owned_by, &hcuser.Login_token, &hcuser.Email_address)
    	}

		hcuser.Sudo = true
	}

    suser, _ := user.Lookup(hcuser.System_username)
    hcuser.HomeDir = suser.HomeDir



    if strings.Contains(hcuser.Privileges, permission) || permission == "any" || strings.Contains(hcuser.Privileges, "all") {
	    return hcuser, true
    }
    
    // fail out with death.
    return hcuser, false
}


func Sudo(username string, ctx *macaron.Context) {
    ctx.SetParams("sudo", username)
}


func ChkPaternity(owner string, child string) (bool) {
    db, _ := MySQL();
	defer db.Close()
	
	
    data := make(map[string]map[string]string)
    data = Getusers(owner, data, db)
    
    if data[child] == nil{
        return false
    }
    
    return true
}

func Getusers(username string, data map[string]map[string]string, db *sql.DB) (map[string]map[string]string){
    stmt, _ := db.Prepare("SELECT * from hostcontrol_users WHERE owned_by = ? and system_username != owned_by")
    rows, _ := stmt.Query(username)
    stmt.Close()
    
    
    
    for rows.Next() {
        var hostcontrol_id string
        var system_username string
        var privileges string
        var owned_by string
        var login_token string
        var email_address string
        
        rows.Scan(&hostcontrol_id,&system_username,&privileges,&owned_by,&login_token,&email_address)
        
        data[system_username] = make(map[string]string)
        data[system_username]["hostcontrol_id"] = hostcontrol_id
        data[system_username]["system_username"] = system_username
        data[system_username]["privileges"] = privileges
        data[system_username]["owned_by"] = owned_by
        data[system_username]["login_token"] = login_token
        data[system_username]["email_address"] = email_address
        
        data = Getusers(system_username, data, db)
    
    }
    
    return data
}

// true means the user has access to the file/dir
func ChkPerms(selected_object string, uid int, gid int) bool {
	if uid == 0 && gid == 0 {
		return true
	}
        fs_object, err := os.Stat(selected_object)
        if err != nil {
                return false
        }

        fs_object_uid := int(fs_object.Sys().(*syscall.Stat_t).Uid)
        fs_object_gid := int(fs_object.Sys().(*syscall.Stat_t).Gid)

        if uid != fs_object_uid || gid != fs_object_gid {
                return false
        }

        return true
}

func Query(ctx *macaron.Context, name string) string {
    query := string(ctx.Query(name))
    param := string(ctx.Params(name))
    if query != "" {
        return query
    }
    return param
}



package rhel7

import (
//	"fmt"
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
    "os/user"
    "strconv"
    "encoding/json"
)


// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func Ftp(ctx *macaron.Context) (string) {
	return "vsftpd"
}

func FtpList(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "ftpusers")
	
	if ! auth {
		return "not_authorized"
	}
	
    db, _ := util.MySQL();
	defer db.Close()
	
	
    data := make(map[string]map[string]string)
    
    stmt, _ := db.Prepare("SELECT * from hostcontrol_ftpusers WHERE system_username = ?")
    rows, _ := stmt.Query(hcuser.System_username)
    stmt.Close()
    
    
    
    for rows.Next() {
        var ftpuser_id string
        var ftpusername string
        var homedir string
        var system_username string
        
        rows.Scan(&ftpuser_id,&ftpusername,&homedir,&system_username)

        data[ftpuser_id] = make(map[string]string)
        data[ftpuser_id]["id"] = ftpuser_id
        data[ftpuser_id]["username"] = ftpusername
        data[ftpuser_id]["homedir"] = homedir
    }
    


    
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}

	return string(output)

}

func AddFtpUser(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "ftpusers")
	if ! auth {
		return "not_authorized"
	}

	suser, err := user.Lookup(hcuser.System_username)

	if err != nil {
		return string(err.Error())
	}

	username := util.Query(ctx, "ftpuser")
	if username == "" {
	    return "ftpuser_required"
	}
	password := util.Query(ctx, "password")
	if password == "" {
	    return "password_required"
	}
	homedir := util.Query(ctx, "homedir")
	if homedir == "" {
	    return "homedir_required"
	}

    username = hcuser.System_username + "_" + username

    // attempt to make homedir as the user
    util.Cmd("su", []string{"-", hcuser.System_username, "-c", "mkdir -p " + homedir})

    // check ownership...
    uid, _ := strconv.Atoi(suser.Uid)
    gid, _ := strconv.Atoi(suser.Gid)
    if ! util.ChkPerms(homedir, uid, gid) {
        return "invalid_homedir"
    }

	db, _ := util.MySQL()
	defer db.Close()
	
	
	// add the user
	// useradd {username} -g {gid} -u {uid} -s /sbin/nologin -o
	util.Cmd("useradd", []string{username,"-d", homedir, "-g", suser.Gid,"-u", suser.Uid, "-s", "/sbin/nologin", "-o"})

	// make sure user was added
    _, lookup_err2 := user.Lookup(username)
    if lookup_err2 != nil {
        return "unable_to_create"
    }
    
    // set the password
	util.Bash("echo " + util.SHSanitize(password) + " | passwd " + util.SHSanitize(username) + " --stdin")
    
	
	// add the user
	istmt, _ := db.Prepare("insert hostcontrol_ftpusers set ftpuser_id=null, ftpusername=?, homedir=?, system_username=?")

	istmt.Exec(username,homedir,hcuser.System_username)
	istmt.Close()


	return "success"
}


func FtpDeleteuser(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "ftpusers")
	if ! auth {
		return "not_authorized"
	}

	username := util.Query(ctx, "ftpuser")

    if username == "" || username == "root" {
        return "ftpuser_required"
    }

	db, _ := util.MySQL()
	defer db.Close()

    // check if user owns domain
    dstmt, _ := db.Prepare("SELECT * FROM `hostcontrol_ftpusers` WHERE `ftpusername`=? and `system_username`=?")
    row1, _ := dstmt.Query(username, hcuser.System_username)
    defer dstmt.Close()
    if ! row1.Next(){
        return "user_not_found"
    }
	
	
	// remove the user
    stmt, _ := db.Prepare("delete from hostcontrol_ftpusers where ftpusername=? and system_username=?")
    stmt.Exec(username, hcuser.System_username)
    stmt.Close()

	// delete the user and homedir
	util.Cmd("userdel", []string{username,"-f"})

	// make sure user was delete
    _, lookup_err2 := user.Lookup(username)
    if lookup_err2 == nil {
        return "failed_to_delete_user"
    }

	return "success"
}



func FtpEditUser(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "ftpusers")
	if ! auth {
		return "not_authorized"
	}

	username := util.Query(ctx, "username")
	password := util.Query(ctx, "password")
	


	db, _ := util.MySQL()
	defer db.Close()

    // check if user owns domain
    dstmt, _ := db.Prepare("SELECT * FROM `hostcontrol_ftpusers` WHERE `ftpusername`=? and `system_username`=?")
    row1, _ := dstmt.Query(username, hcuser.System_username)
    defer dstmt.Close()
    if ! row1.Next(){
        return "user_not_found"
    }
	
	
    // set the password
	util.Bash("echo " + util.SHSanitize(password) + " | passwd " + util.SHSanitize(username) + " --stdin")
    

	return "success"
}

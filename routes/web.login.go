package routes

import (
	"strconv"
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"
	"strings"
)


func init(){
	route("/", login, "GET")
	route("/login", login_post, "POST")
	route("/logout", logout, "GET")
	route("/dashboard", dashboard, "BOTH")
}


func dashboard(ctx *macaron.Context) (string){
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/dashboard.tpl")

	hostname := string(ctx.Req.Header.Get("X-FORWARDED-HOST"))
	if hostname == "" {
		hostname = string(ctx.Req.Host)
	}
	hostname = strings.Split(hostname, ":")[0]
	tpl.Assign("console_url", "https://"+hostname+"/shellinabox")

	tpl.Parse("dashboard")

    if (strings.Contains(hcuser.Privileges, "websites") || strings.Contains(hcuser.Privileges, "all")) && hcuser.System_username != "root" {
        tpl.Parse("dashboard/websitesbtn");
    }
    if strings.Contains(hcuser.Privileges, "databases") || strings.Contains(hcuser.Privileges, "all") {
        tpl.Parse("dashboard/databasesbtn");
    }
    if strings.Contains(hcuser.Privileges, "dns") || strings.Contains(hcuser.Privileges, "all") {
        tpl.Parse("dashboard/dnsbtn");
    }
    if (strings.Contains(hcuser.Privileges, "mail") || strings.Contains(hcuser.Privileges, "all")) && hcuser.System_username != "root" {
        tpl.Parse("dashboard/mailbtn");
    }
    if (strings.Contains(hcuser.Privileges, "ftpusers") || strings.Contains(hcuser.Privileges, "all")) && hcuser.System_username != "root" {
        tpl.Parse("dashboard/ftpusersbtn");
    }
    if strings.Contains(hcuser.Privileges, "all") {
        tpl.Parse("dashboard/firewallbtn");
    }
    if strings.Contains(hcuser.Privileges, "all") {
        tpl.Parse("dashboard/servicesbtn");
    }
    if strings.Contains(hcuser.Privileges, "sysusers") || strings.Contains(hcuser.Privileges, "all") {
        tpl.Parse("dashboard/usersbtn");
    }
	return header(ctx) + tpl.Out() + footer(ctx)
}

func logout(ctx *macaron.Context) (string) {
	var tpl vision.New
	tpl.TemplateFile("template/login.tpl")

	user, auth := util.Auth(ctx, "any")

	if user.Sudo {
		ctx.SetCookie("sudo", "", -1)
		set_error("No longer logged in as " + user.System_username + ".", ctx)
		ctx.Redirect("/dashboard", 302)
		return "success"
	}

	if auth {
		new_token := util.MkToken()
		db, _ := util.MySQL()	
		defer db.Close()

		ustmt, _ := db.Prepare("update hostcontrol_users set login_token=? where system_username=?")
		ustmt.Exec(new_token, user.System_username)
		ustmt.Close()
	}


	ctx.SetCookie("hostcontrol_id", "", -1)
	ctx.SetCookie("login_token", "", -1)

	tpl.Parse("login")
	tpl.Parse("login/logged_out")
	return tpl.Out()
	
}

func login(ctx *macaron.Context) (string) {
	var tpl vision.New
	tpl.TemplateFile("template/login.tpl")
	tpl.Assign("x", "y")
	tpl.Parse("login")
	return tpl.Out()
	
}

func login_post(ctx *macaron.Context) (string) {
	db, err := util.MySQL()
	defer db.Close()
	if err != nil {
		return "Problem opening MySQL"
	}

	new_token := util.MkToken()

	username := util.Query(ctx, "username")
	password := util.Query(ctx, "password")
	rememberme := util.Query(ctx, "rememberme")

	login_failed := false

	if chklogin(username,password) {
		stmt, _ := db.Prepare("SELECT * from hostcontrol_users WHERE system_username = ?")
		rows, _ := stmt.Query(username)
		stmt.Close()

		var hostcontrol_id int
		var system_username string
		var privileges string
		var owned_by string
		var login_token string
		var email_address string

		// check if we have a row returned...
		if rows.Next() {
			rows.Scan(&hostcontrol_id, &system_username, &privileges, &owned_by, &login_token, &email_address)
			ustmt, _ := db.Prepare("update hostcontrol_users set login_token=? where system_username=?")
			ustmt.Exec(new_token, username)
			ustmt.Close()

		// insert root if login worked and he doesn't exist!
		}else if username == "root" {
			istmt, _ := db.Prepare("insert hostcontrol_users set hostcontrol_id=null, system_username=?, privileges=?, owned_by=?, login_token=?, email_address=?")
			istmt.Exec("root","all","root",new_token,"")
			istmt.Close()

		// fallback to failure.
		}else{
			login_failed = true
		}

		if ! login_failed {
			// set cookies
			if rememberme == "checked" {
				ctx.SetCookie("hostcontrol_id", strconv.Itoa(hostcontrol_id), 864000)
				ctx.SetCookie("login_token", new_token, 864000)
				ctx.SetCookie("sudo", "", 864000)
			}else{
				ctx.SetCookie("hostcontrol_id", strconv.Itoa(hostcontrol_id), 0)
				ctx.SetCookie("login_token", new_token, 0)
				ctx.SetCookie("sudo", "", 0)
			}
		
			// send to dashboard
			ctx.Redirect("/dashboard", 302)
			return "Redirecting to the dashboard. Click <a href=\"/dashboard\">here</a> if you are not redirected."
		}
	}else{
		login_failed = true
	}

	var tpl vision.New
	tpl.TemplateFile("template/login.tpl")

	tpl.Parse("login")

	if login_failed {
		tpl.Parse("login/fail")
	}

	return tpl.Out()
	
}

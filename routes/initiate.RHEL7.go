package routes

import (
//	"fmt"
//	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/routes/rhel7"

	"io/ioutil"
	"regexp"
	"strconv"
)

func init(){
	// Make sure we're on RHEL7. If we aren't, do not add RHEL7 routes to web api.
	if !chk_RHEL7() {
		return
	}

	// Add routes to our API
	//    Path				Function		Method
	//------------				--------		------
	route("/api/distro",			    rhel7.Distro,	      "BOTH")
	route("/api/install",			    rhel7.Install,	      "BOTH")
            
	route("/api/login",		    	    rhel7.Distro,	      "BOTH")
	route("/api/logout",			    rhel7.Distro,	      "BOTH")
	route("/api/authcheck",			    rhel7.Distro,	      "BOTH")
            
	route("/api/users/list",		    rhel7.Listusers,   "BOTH")
	route("/api/users/add",			    rhel7.Adduser,	      "BOTH")
	route("/api/users/delete",		    rhel7.Deleteuser,      "BOTH")
	
	route("/api/ftp",                   rhel7.Ftp,              "BOTH")
	route("/api/ftpusers/list",              rhel7.FtpList,              "BOTH")
	route("/api/ftpusers/add",               rhel7.AddFtpUser,              "BOTH")
	route("/api/ftpusers/edit",              rhel7.FtpEditUser,              "BOTH")
	route("/api/ftpusers/delete",            rhel7.FtpDeleteuser,              "BOTH")
            
	route("/api/dns",		    	    rhel7.Dns,  	      "BOTH")
	route("/api/dns/list",			    rhel7.DnsList,	          "BOTH")
	route("/api/dns/domain/add",		rhel7.DnsAddDomain,  "BOTH")
	route("/api/dns/domain/delete",	    rhel7.DnsDeleteDomain,  "BOTH")
	route("/api/dns/record/edit",	    rhel7.DnsEditRecord,	          "BOTH")
	route("/api/dns/record/add",	    rhel7.DnsAddRecord,	          "BOTH")
	route("/api/dns/record/delete",	    rhel7.DnsDeleteRecord,     "BOTH")
            
	route("/api/mail",		    	    rhel7.Mail,		   "BOTH")
	route("/api/mail/list",			    rhel7.MailList,	      "BOTH")
	route("/api/mail/domain/add",	    rhel7.MailAddDomain,	          "BOTH")
	route("/api/mail/domain/delete",    rhel7.MailDeleteDomain,	            "BOTH")
	route("/api/mail/users/add",	    rhel7.MailAddUser,	          "BOTH")
	route("/api/mail/users/delete",     rhel7.MailUserDelete,	            "BOTH")
	route("/api/mail/users/edit",	    rhel7.MailEditUser,		        "BOTH")
        
	route("/api/sql",		        	rhel7.Sql,	            "BOTH")
	
	route("/api/sql/databases/list",	rhel7.SqlDatabasesList,	 "BOTH")
	route("/api/sql/databases/add",		rhel7.SqlDatabasesAdd,	 "BOTH")
	route("/api/sql/databases/delete",	rhel7.SqlDatabasesDelete,"BOTH")
        
	route("/api/sql/users/list",		rhel7.SqlUsersList,     "BOTH")
	route("/api/sql/users/add",	    	rhel7.SqlUsersAdd,       "BOTH")
	route("/api/sql/users/delete",		rhel7.SqlUsersDelete,    "BOTH")
	route("/api/sql/users/edit",		rhel7.SqlUsersEdit,      "BOTH")
	        
	route("/api/sql/grants/list",		rhel7.SqlGrantsList,    "BOTH")
	route("/api/sql/grants/add",		rhel7.SqlGrantsAdd,	    "BOTH")
	route("/api/sql/grants/delete",		rhel7.SqlGrantsDelete,  "BOTH")

	route("/api/web",           		rhel7.Web,      	    "BOTH")
	route("/api/web/domain/list",		rhel7.ListWebsites,	    "BOTH")
	route("/api/web/domain/add",		rhel7.AddWebsite,   	"BOTH")
	route("/api/web/domain/delete",		rhel7.DeleteWebsite,	"BOTH")
	route("/api/web/domain/sslmanage",		rhel7.ManageWebsiteSSL,	"BOTH")

	// Add our dynamic functions below.
	// setup our login method
	chklogin = rhel7.ChkLogin
	// setup check password method
	chpassword = rhel7.ChPassword
	
	// backmount the API calls so we can use API() in the distro package.
    rhel7.Route_function = route_function
	
}


func chk_RHEL7() bool {
	fbuffer, err := ioutil.ReadFile("/etc/redhat-release")

	if err != nil {
		return false // not RHEL 7 variant
	}
	
	release_file := string(fbuffer)

	ex := regexp.MustCompile("release ([0-9]+)")
	raw_release := ex.FindStringSubmatch(release_file)

	release, _ := strconv.Atoi(raw_release[1])
	
	if release != 7 {
		return false // not RHEL 7 variant
	}
	
	return true
}

package rhel7

import (
//	"fmt"
	"gopkg.in/macaron.v1"
	"github.com/protosam/hostcontrol/util"
	
	"os"
	"os/user"
	"io/ioutil"
	"path"
	"strings"
	"strconv"
	
	"encoding/json"
	"crypto/tls"
)


func ManageWebsiteSSL(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "websites")
	if ! auth {
		return "not_authorized"
	}

    db, err := util.MySQL();
    if err != nil {
        return string(err.Error())
    }
    defer db.Close()
    
    vhost_id := util.Query(ctx, "vhost_id")
    
    enablessl := util.Query(ctx, "enablessl")
    
    crt_data := strings.Trim(util.Query(ctx, "crt_data"), " ")
    crtca_data := strings.Trim(util.Query(ctx, "crtca_data"), " ")
    key_data := strings.Trim(util.Query(ctx, "key_data"), " ")
    
    if enablessl != "Y" {
        enablessl = "N"
    }

    
    
    stmt, _ := db.Prepare("SELECT * from website_vhosts WHERE vhost_id = ? and system_username=?")
    rows, _ := stmt.Query(vhost_id, hcuser.System_username)
    stmt.Close()
    
    
    if rows.Next() {
        var vhost_id string
        var system_username string
        var domain string
        var documentroot string
        var ipaddr string
        var ssl_enabled string
        var ssl_certificate string
        var ssl_key string
        var ssl_ca_certificate string
        
        rows.Scan(&vhost_id, &system_username, &domain, &documentroot, &ipaddr, &ssl_enabled, &ssl_certificate, &ssl_key, &ssl_ca_certificate)

    
        crt_write := ioutil.WriteFile("/etc/pki/tls/certs/"+domain+".crt.tmp", []byte(crt_data), 0644)
    	if crt_write != nil {
    		return string(crt_write.Error())
    	}
        crtca_write := ioutil.WriteFile("/etc/pki/tls/certs/"+domain+".ca.crt.tmp", []byte(crtca_data), 0644)
    	if crtca_write != nil {
    		return string(crtca_write.Error())
    	}
        key_write := ioutil.WriteFile("/etc/pki/tls/private/"+domain+".key.tmp", []byte(key_data), 0644)
    	if key_write != nil {
    		return string(key_write.Error())
    	}

    	_, crt_err := tls.LoadX509KeyPair("/etc/pki/tls/certs/"+domain+".crt.tmp", "/etc/pki/tls/private/"+domain+".key.tmp")
    	if crt_err != nil {
        	os.RemoveAll("/etc/pki/tls/certs/"+domain+".crt.tmp")
        	os.RemoveAll("/etc/pki/tls/certs/"+domain+".ca.crt.tmp")
        	os.RemoveAll("/etc/pki/tls/private/"+domain+".key.tmp")
    		return "certificate_key_pair_failed " + string(crt_err.Error())
    	}
    	
    	os.RemoveAll("/etc/pki/tls/certs/"+domain+".crt")
    	os.RemoveAll("/etc/pki/tls/certs/"+domain+".ca.crt")
    	os.RemoveAll("/etc/pki/tls/private/"+domain+".key")
    	
    	os.Rename("/etc/pki/tls/certs/"+domain+".crt.tmp", "/etc/pki/tls/certs/"+domain+".crt")
    	os.Rename("/etc/pki/tls/certs/"+domain+".ca.crt.tmp", "/etc/pki/tls/certs/"+domain+".ca.crt")
    	os.Rename("/etc/pki/tls/private/"+domain+".key.tmp", "/etc/pki/tls/private/"+domain+".key")
    	
    	rawvhostconf, _ := ioutil.ReadFile("common/src/rhel7/httpd/vhost.ssl.conf")
    	vhost_data := string(rawvhostconf)
    	
    	vhost_data = strings.Replace(vhost_data, "__IPADDR__", "*", -1)
    	vhost_data = strings.Replace(vhost_data, "__HOSTNAME__", domain, -1)
    	vhost_data = strings.Replace(vhost_data, "__USERNAME__", hcuser.System_username, -1)
    	vhost_data = strings.Replace(vhost_data, "__DOCUMENTROOT__", documentroot, -1)
    	vdat := []byte(vhost_data)
    	
        write_err := ioutil.WriteFile("/etc/httpd/vhosts.d/" + domain + ".ssl.conf", vdat, 0644)
    	if write_err != nil {
    		return string(write_err.Error())
    	}

        xstmt, err := db.Prepare("update website_vhosts set ssl_enabled=?, ssl_certificate=?, ssl_key=?, ssl_ca_certificate=? where vhost_id=?")
        if err != nil {
    	    return "failed_to_update_record"
    	}
    	
    	_, xerr := xstmt.Exec(enablessl, crt_data, key_data, crtca_data, vhost_id)
    	xstmt.Close()
    
    	if xerr != nil {
    	    return "failed_to_update_record"
    	}
    }else{
        return "domain_not_found"
    }

    util.Bash("systemctl reload httpd")
	return "success"
}


func ListWebsites(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "websites")
	
	if ! auth {
		return "not_authorized"
	}
	
    db, err := util.MySQL();
    defer db.Close()
    
    stmt, _ := db.Prepare("SELECT * from website_vhosts WHERE system_username = ?")
    rows, _ := stmt.Query(hcuser.System_username)
    stmt.Close()
    
    
    data := make(map[string]map[string]string)
    
    for rows.Next() {
        var vhost_id string
        var system_username string
        var domain string
        var documentroot string
        var ipaddr string
        var ssl_enabled string
        var ssl_certificate string
        var ssl_key string
        var ssl_ca_certificate string
        
        rows.Scan(&vhost_id, &system_username, &domain, &documentroot, &ipaddr, &ssl_enabled, &ssl_certificate, &ssl_key, &ssl_ca_certificate)
        
        data[domain] = make(map[string]string)
        data[domain]["vhost_id"] = vhost_id
        data[domain]["system_username"] = system_username
        data[domain]["domain"] = domain
        data[domain]["documentroot"] = documentroot
        data[domain]["ipaddr"] = ipaddr
        data[domain]["ssl_enabled"] = ssl_enabled
        data[domain]["ssl_certificate"] = ssl_certificate
        data[domain]["ssl_key"] = ssl_key
        data[domain]["ssl_ca_certificate"] = ssl_ca_certificate
    }
    
    output, err := json.Marshal(data)
	if err != nil {
		return "json_out_failed: " + string(err.Error())
	}

	return string(output)
}

func DeleteWebsite(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "websites")
	if ! auth {
		return "not_authorized"
	}

    db, err := util.MySQL();
    if err != nil {
        return string(err.Error())
    }
    defer db.Close()
    
    vhost_id := util.Query(ctx, "vhost_id")
    
    stmt, _ := db.Prepare("SELECT * from website_vhosts WHERE vhost_id = ? and system_username=?")
    rows, _ := stmt.Query(vhost_id, hcuser.System_username)
    stmt.Close()
    
    
    if rows.Next() {
        var vhost_id string
        var system_username string
        var domain string
        var documentroot string
        var ipaddr string
        var ssl_enabled string
        var ssl_certificate string
        var ssl_key string
        var ssl_ca_certificate string
        
        rows.Scan(&vhost_id, &system_username, &domain, &documentroot, &ipaddr, &ssl_enabled, &ssl_certificate, &ssl_key, &ssl_ca_certificate)
        
        os.RemoveAll("/var/log/httpd/" + hcuser.System_username + "/" + domain + "-error_log")
        os.RemoveAll("/var/log/httpd/" + hcuser.System_username + "/" + domain + "-access_log")
        os.RemoveAll("/var/log/httpd/" + hcuser.System_username + "/" + domain + "-ssl-error_log")
        os.RemoveAll("/var/log/httpd/" + hcuser.System_username + "/" + domain + "-ssl-access_log")
        os.RemoveAll("/etc/pki/tls/certs/"+domain+".crt")
        os.RemoveAll("/etc/pki/tls/certs/"+domain+".ca.crt")
        os.RemoveAll("/etc/pki/tls/private/"+domain+".key")
        os.RemoveAll("/etc/httpd/vhosts.d/" + domain + ".conf")
        os.RemoveAll("/etc/httpd/vhosts.d/" + domain + ".ssl.conf")
        stmt, _ = db.Prepare("delete from website_vhosts where vhost_id=?")
        stmt.Exec(vhost_id)
        stmt.Close()

    }else{
        return "domain_not_found"
    }

    util.Bash("systemctl reload httpd")
	return "success"
}

func AddWebsite(ctx *macaron.Context) (string) {
	hcuser, auth := util.Auth(ctx, "websites")
	if ! auth {
		return "not_authorized"
	}
	
    db, err := util.MySQL();
    if err != nil {
        return string(err.Error())
    }
    
    defer db.Close()
    
	domainname := util.Query(ctx, "domainname")
    if domainname == "" {
        return "invalid_domainname"
    }

    // check if website is taken already
    stmt, _ := db.Prepare("SELECT * from website_vhosts WHERE domain = ?")
    rows, _ := stmt.Query(domainname)
    stmt.Close()
    
    if rows.Next() {
		return "not_authorized"
    }
    
    suser, err := user.Lookup(hcuser.System_username)
    
	documentroot := path.Clean("/www/" + domainname)
	documentroot = path.Clean(suser.HomeDir + "/" + documentroot)
	documentroot_base := path.Clean(suser.HomeDir + "/www")
    
    uid, err := strconv.Atoi(suser.Uid)
	if err != nil {
		return string(err.Error())
	}


	gid, err := strconv.Atoi(suser.Gid)
	if err != nil {
		return string(err.Error())
	}
    
    os.Mkdir("/var/log/httpd/" + hcuser.System_username, 0755)
    
	os.MkdirAll(documentroot, 0755)
	os.Chown(documentroot, uid, gid)
	os.Chown(documentroot_base, uid, gid)
	
	rawvhostconf, _ := ioutil.ReadFile("common/src/rhel7/httpd/vhost.conf")
	vhost_data := string(rawvhostconf)
	
	vhost_data = strings.Replace(vhost_data, "__HOSTNAME__", domainname, -1)
	vhost_data = strings.Replace(vhost_data, "__USERNAME__", hcuser.System_username, -1)
	vhost_data = strings.Replace(vhost_data, "__DOCUMENTROOT__", documentroot, -1)
	vdat := []byte(vhost_data)
	
    write_err := ioutil.WriteFile("/etc/httpd/vhosts.d/" + domainname + ".conf", vdat, 0644)

	if write_err != nil {
		return string(write_err.Error())
	}

	istmt, _ := db.Prepare("INSERT INTO `hostcontrol`.`website_vhosts` set `vhost_id`=null, `system_username`=?, `domain`=?, `documentroot`=?, `ipaddr`=?, `ssl_enabled`='N', `ssl_certificate`='', `ssl_key`='', `ssl_ca_certificate`=''")
	istmt.Exec(hcuser.System_username,domainname,documentroot,"*")
	istmt.Close()
	
	util.Bash("systemctl reload httpd")

    return "success"
}
// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func Web(ctx *macaron.Context) (string) {
	return "httpd"
}



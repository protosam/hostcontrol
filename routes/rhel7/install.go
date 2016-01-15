package rhel7

import (
//	"fmt"
	"gopkg.in/macaron.v1"
)






func Install(ctx *macaron.Context) (string) {
	//username := util.Query(ctx, "username")
	// will need to detect reading CLI args in.

	runscript("assets/install.prep.sh")
	runscript("assets/install.mysql.sh")
	runscript("assets/install.httpd.sh")
	runscript("assets/install.mail.sh")
	runscript("assets/install.dns.sh")
	runscript("assets/install.ftp.sh")
	return "success"
}


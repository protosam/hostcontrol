package main

import (
	"fmt"
	"github.com/protosam/hostcontrol/routes"
	"github.com/protosam/hostcontrol/util"
	"gopkg.in/macaron.v1"
	"net/http"

	"os"
	//	"github.com/protosam/hostcontrol/services/RHEL7/users"

	"strings"
)

// This has our server variable stuff: https://godoc.org/github.com/Unknwon/macaron#Context
var Webserver *macaron.Macaron

func main() {

	_, err := os.Stat("./dev_run.sh")
	if err != nil {
		os.Chdir("/opt/hostcontrol/")
	}

	config, err := util.ReadConfig("settings.cfg")

	if err != nil {
		fmt.Println(err)
		return
	}


	// for CLI usage
	if len(os.Args) > 1 {
		args := strings.Split(os.Args[1], "?")
		fmt.Println(routes.CLI(args[0]))
		return
	}

	Webserver = macaron.Classic()

	routes.StartApi(Webserver)

	// Serve up neccessary static content from the root of site
	Webserver.Use(macaron.Static("www"))
	Webserver.Use(macaron.Static("/opt/hostcontrol/www"))

	fmt.Println("Running webserver on: " + config.Webserver.Bind + ":" + config.Webserver.Port)
	http.ListenAndServe(config.Webserver.Bind+":"+config.Webserver.Port, Webserver)
}

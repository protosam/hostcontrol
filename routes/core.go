package routes

import (
	"gopkg.in/macaron.v1"
	"fmt"
	"strings"
	"github.com/protosam/hostcontrol/util"
	//"net/url"
//	"github.com/protosam/hostcontrol/services/dns"
//	"github.com/protosam/hostcontrol/services/mail"
//	"github.com/protosam/hostcontrol/services/domains"
//	"github.com/protosam/hostcontrol/services/users"
//	"github.com/protosam/hostcontrol/services/sql"
)

var Webserver *macaron.Macaron
var routes = make(map[string]func())


var route_function = make(map[string]util.API_function)
var route_method = make(map[string]string)


func StartApi(webserver *macaron.Macaron){
	// Setup Webserver variable
	Webserver = webserver

	// Let the user know we're binding the API services.
	fmt.Println("Starting api services...")

	// We will add all of our routes arbitrarily.
	for name := range route_function {
		fmt.Println("Adding route " + name)

		if route_method[name] == "GET" {
			Webserver.Get(name, route_function[name])
		} else if route_method[name] == "POST" {
			Webserver.Post(name, route_function[name])
		} else {
			Webserver.Get(name, route_function[name])
			Webserver.Post(name, route_function[name])
		}
	}

}


func CLI(routename string) string {
	return route_function[routename](nil)
}



func API(rawuri string, ctx *macaron.Context) (string) {
    querystr := strings.Split(rawuri, "?")
    routename := strings.Split(rawuri, "?")[0]
    if len(querystr) > 1 {
        
        params := strings.Split(querystr[1], "&")
        for _, param := range params {
            par := strings.Split(param, "=")
            if len(par) > 1 {
                ctx.SetParams(par[0], par[1])
            }
            
        }
    }
    
    ret := route_function[routename](ctx)
	return ret
}

// Function for adding new route dynamically
func route(routename string, fn util.API_function, method string){
	route_method[routename] = method
	route_function[routename] = fn;

}

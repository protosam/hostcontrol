////////////////////////////////////////////////////////////////////////////////
// This file is absolutely neccesary so that we can use the API() command in 
// our API... IT IS REQUIRED. COPY IT.

package rhel7

import (
	"gopkg.in/macaron.v1"
	"strings"
	"github.com/protosam/hostcontrol/util"
)

var Route_function = make(map[string]util.API_function)

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
    ret := Route_function[routename](ctx)
	return ret
}

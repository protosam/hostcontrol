package rhel7

import (
	"gopkg.in/macaron.v1"
)



// This will return RHEL7 for the server API test. Note that all functions need to be prefixed with DISTRO TAG.
func Sql(ctx *macaron.Context) (string) {
	return "mariadb"
}


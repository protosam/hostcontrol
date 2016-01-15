package rhel7

import (
	"github.com/protosam/hostcontrol/util"
)


func runscript(asset_path string) {
	data, _ := Asset(asset_path)
	bash_script := string(data)
	util.Bash(bash_script)
}


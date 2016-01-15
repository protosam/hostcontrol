package util

import (
	"gopkg.in/gcfg.v1"
)


type Config struct {
	Hostcontrol struct {
		License string
	}
	Webserver struct {
		Bind string
		Port string
	}
	MySQL struct {
		Dbname string
		Host string
		User string
		Pass string
		Port string
		Socket string
	}
}

func ReadConfig(config_file string) (Config, error) {
	var cfg Config
	
	err := gcfg.ReadFileInto(&cfg, config_file)
	
	return cfg, err
}



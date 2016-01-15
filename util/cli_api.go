package util

import (
	"strings"
	"os"
)

func GetCLI(key string) map[string]string {
	variables := make(map[string]string)

	args := strings.Split(os.Args[1], "?")
	pre_map := strings.Split(args[1], "&")
	for i := range pre_map {
		raw := strings.Split(pre_map[i], "=")
		variables[raw[0]] = raw[1]
	}

	return variables
}

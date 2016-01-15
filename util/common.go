package util

import (
	"fmt"
	"os/exec"
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func Bash(command string) string {

	args := []string{"-c", command}
	cmdOut, err := exec.Command("bash", args...).Output();
	cmd := string(cmdOut)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Print(cmd)

	return cmd
}

func SHSanitize(str string) string {
	str = strings.Replace(str, "'", "'\"'\"'", -1)
	return "'" + str + "'"
}

func Cmd(cmd string, args []string) (string, error) {
	cmdOut, err := exec.Command(cmd, args...).Output();
	return string(cmdOut), err
}


func MD5SUM(path string) string {
	return Bash("find " + path + " -type f -exec md5sum {} \\; | sort -k 34 | md5sum | cut -d ' ' -f 1")
}

func SurfGet(siteurl string) string {

	resp, err := http.Get(siteurl)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	
	return string(body)


}

func SurfPost(siteurl string, postdata string) string {

	resp, err := http.Post(siteurl, "application/x-www-form-urlencoded", bytes.NewBufferString(postdata))

	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return string(body)

}


func FileLockCheck(str string, list []string) bool {
    for i :=0; i < len(list); i++ {
		if list[i] == str || strings.HasPrefix(str, list[i]) {
			return true
		}
	}
	return false
}

func IndexOfString(str string, list []string) int {
    for i :=0; i < len(list); i++ {
		if list[i] == str {
			return i
		}
	}
	return -1
}


func RemoveInArray(str string, list []string) []string {
	i := IndexOfString(str, list)
	for i > -1 {
		list = append(list[:i], list[i+1:]...)
		i = IndexOfString(str, list)
	}
	return list
}



func IsInterface(ip_addr string) bool {
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
			}
			if ip_addr == ip.String() {
				return true
			}
		}
	}
	return false
}



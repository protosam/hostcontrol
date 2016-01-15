package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"crypto/rand"
	"strings"
)

// Example taken from here: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html
func MySQL() (*sql.DB, error) {
	var db *sql.DB
	var err error

	config, cerr := ReadConfig("settings.cfg")

	if cerr != nil {
		fmt.Println(cerr)
		return nil, cerr
	}

	if config.MySQL.Socket != "" {
		db, err = sql.Open("mysql", config.MySQL.User+":"+config.MySQL.Pass+"@unix("+config.MySQL.Socket + ")/"+config.MySQL.Dbname)
	}else{
		db, err = sql.Open("mysql", config.MySQL.User+":"+config.MySQL.Pass+"@tcp("+config.MySQL.Host + ":" + config.MySQL.Port +")/"+config.MySQL.Dbname)
	}


	if err != nil {
		fmt.Println(err)
	}

	return db, err
}

func LastResortSanitize(str string) (string) {
    // fucking escape artists
    str = strings.Replace(str, "\\", "\\\\'", -1)
    // fucking quotes...
    str = strings.Replace(str, "'", "\\'", -1)
    str = strings.Replace(str, "\"", "\\\"", -1)
    
    return str
}

func MkToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

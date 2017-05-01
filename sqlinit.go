package mysqlinit

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"net"
	"errors"
)

var (
	dbConfigPath = "configuration/DBCredentials.json"
	ipAddress = "127.0.0.1"
	portNumber = "3306"
	databaseName string
)
func SetDBConfigPath(path string) {
	dbConfigPath = path
}

func SetIPAddr(addr string) error {
	ip := net.ParseIP(addr)
	if ip == nil {
		return errors.New("malformed IP address")
	}

	ipAddress = addr
	return nil
}

func SetPort(port int) error {
	if port < 0 {
		return errors.New("negative port number")
	}

	if port > 65535 {
		return errors.New("port number exceeds maximum value")
	}

	portNumber = string(port)
	return nil
}

func SetDatabaseName(name string) {
	databaseName = name
}

func InitCnxn() (*sql.DB, error) {
	if databaseName == "" {
		return nil, errors.New("database name not set")
	}

	file, err := ioutil.ReadFile(dbConfigPath)
	if err != nil {
		return nil, err
	}

	var dbCredentials struct{
		Username string
		Password string
	}

	err = json.Unmarshal(file, &dbCredentials)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		ipAddress,
		portNumber,
		dbCredentials.Username,
		dbCredentials.Password,
		databaseName,
	)

	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}


package mysqlclient

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func NewMysqlClient(
	host,
	port,
	user,
	pass,
	dbName string,
	maxConn,
	maxIdleConn int,
) (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Etc/UTC")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil && viper.GetBool("DEBUG") {
		return nil, err
	}

	return db, nil
}

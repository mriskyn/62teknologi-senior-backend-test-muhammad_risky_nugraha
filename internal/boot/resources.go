package boot

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	_mysql "62teknologi-senior-backend-test-muhammad_risky_nugraha/pkg/db/mysqlclient"
)

var (
	MainDBConn      *gorm.DB
)

func init() {
	// set config file .env on root directory
	viper.AddConfigPath(`./`)
	viper.SetConfigFile(`.env`)

	fmt.Println("VIPER ENV", viper.ReadInConfig())
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("DEBUG") {
		fmt.Println("Service RUN on DEBUG mode")
	}

	// init sql client
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPass := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbMaxConn := viper.GetInt("DB_MAX_CONN")
	dbMaxIdleConn := viper.GetInt("DB_MAX_IDLE_CONN")
	MainDBConn, err = _mysql.NewMysqlClient(dbHost, dbPort, dbUser, dbPass, dbName, dbMaxConn, dbMaxIdleConn)
	if err != nil {
		log.Fatalf("MainDBConn.Init: %s", err)
	}
}


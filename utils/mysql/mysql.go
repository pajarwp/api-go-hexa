package mysql

import (
	"api-go-hexa/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase(config *config.MySQLConfig) *sql.DB {
	// init database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Name)
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}
	return dbConn
}

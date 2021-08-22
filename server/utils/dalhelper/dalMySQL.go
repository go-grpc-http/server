package dalhelper

import (
	"gin-solid-template/util/confighelper"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Hold a single global connection (pooling provided by sql driver)
var sqlConnection *sqlx.DB
var connectionError error
var sqlOnce sync.Once

//GetMySQLConnection creates MySQLConnection.
//Returns connection if established else error
func GetMySQLConnection() (*sqlx.DB, error) {
	//sqlOnce is used to create singleton object
	sqlOnce.Do(func() {
		// create a connection db(e.g. "postgres", "mysql", or "sqlite3")
		connection, err := sqlx.Open("mysql", confighelper.GetConfig("MySqlDSN"))
		if err != nil {
			connectionError = err
		}
		//set maximum number of idle connections to handle
		connection.SetMaxIdleConns(100)
		//set maximum number of open connections to handle
		connection.SetMaxOpenConns(1000)
		//Connection alive duration
		// duration := 3 * 24 * time.Hour
		duration := 4 * time.Hour
		connection.SetConnMaxLifetime(duration)
		sqlConnection = connection
	})
	return sqlConnection, connectionError
}

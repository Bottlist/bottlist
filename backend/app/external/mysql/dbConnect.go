package mysql

import (
	"fmt"
	"github.com/Bottlist/bottlist/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const driverName = "mysql"

func NewMySQLConnector() *sqlx.DB {
	conf := config.LoadConfig()

	dsn := mysqlConnInfo(*conf.MySQLInfo)
	conn, err := sqlx.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return conn
}

func mysqlConnInfo(mysqlInfo config.MysqlInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		mysqlInfo.MySQLUser,
		mysqlInfo.MySQLPassWord,
		mysqlInfo.MySQLHost,
		mysqlInfo.MySQLPort,
		mysqlInfo.MySQLDBName)

	return dataSourceName
}

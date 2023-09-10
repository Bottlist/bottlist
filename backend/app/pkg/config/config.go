package config

import (
	"os"
)

type appConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MysqlInfo
}

type HTTPInfo struct {
	Addr string
}

type MysqlInfo struct {
	MySQLDBName   string
	MySQLUser     string
	MySQLPassWord string
	MySQLHost     string
	MySQLPort     string
}

func LoadConfig() *appConfig {
	addr := ":" + os.Getenv("PORT")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	mysqlDBName := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	dbInfo := &MysqlInfo{
		MySQLDBName:   mysqlDBName,
		MySQLUser:     mysqlUser,
		MySQLPassWord: mysqlPassword,
		MySQLHost:     mysqlHost,
		MySQLPort:     mysqlPort,
	}

	conf := appConfig{
		MySQLInfo: dbInfo,
		HTTPInfo:  httpInfo,
	}

	return &conf
}

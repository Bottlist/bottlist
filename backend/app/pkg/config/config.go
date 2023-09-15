package config

import (
	"os"
)

type appConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MysqlInfo
	RedisInfo *RedisInfo
	MailInfo  *MailInfo
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

type RedisInfo struct {
	RedisHost string
	RedisPort string
}

type MailInfo struct {
	HostName string
	Port     string
	UserName string
	PassWord string
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

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisInfo := &RedisInfo{
		RedisHost: redisHost,
		RedisPort: redisPort,
	}
	mailHost := os.Getenv("MAIL_HOST")
	mailPort := os.Getenv("MAIL_PORT")
	mailUser := os.Getenv("MAIL_USER")
	mailPass := os.Getenv("MAIL_PASS")

	mailInfo := &MailInfo{
		mailHost,
		mailPort,
		mailUser,
		mailPass,
	}

	conf := appConfig{
		MySQLInfo: dbInfo,
		HTTPInfo:  httpInfo,
		RedisInfo: redisInfo,
		MailInfo:  mailInfo,
	}

	return &conf
}

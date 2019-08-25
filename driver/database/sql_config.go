package database

import "fmt"

type SQLConfig struct {
	SQLName  string
	Port     string
	Host     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

func NewSQLConfig(sqlName string, port string, host string, user string, password string, dbName string, sslMode string) *SQLConfig {
	return &SQLConfig{
		SQLName:  sqlName,
		Port:     port,
		Host:     host,
		User:     user,
		Password: password,
		DBname:   dbName,
		SSLMode:  sslMode,
	}
}

func (c *SQLConfig) GetStr() string {
	return fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Port,
		c.Host,
		c.User,
		c.DBname,
		c.Password,
		c.SSLMode,
	)
}

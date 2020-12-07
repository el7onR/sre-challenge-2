package config

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectionInfo is the info of the mysql
type ConnectionInfo struct {
	Username string
	Password string
	Address  string
	DBName   string
}

func Database(ctx context.Context, c *ConnectionInfo) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=skip-verify&charset=utf8&parseTime=true", c.Username, c.Password, c.Address, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

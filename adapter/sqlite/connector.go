package sqlite

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PoolInterface interface {
	Create(value interface{}) (tx *gorm.DB)
}

func GetConnection() *gorm.DB {
	databaseUrl := viper.GetString("database.url")
	conn, err := gorm.Open(sqlite.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

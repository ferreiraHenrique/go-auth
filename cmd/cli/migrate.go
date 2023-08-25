package main

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	db := sqlite.GetConnection()

	if db.Migrator().HasTable(&domain.Manager{}) {
		db.Migrator().DropTable(&domain.Manager{})
	}
	db.Migrator().CreateTable(&domain.Manager{})
}

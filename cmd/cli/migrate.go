package main

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
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
	sqlite.Migrator()
}

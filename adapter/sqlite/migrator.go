package sqlite

import (
	"fmt"

	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func Migrator() {
	db := GetConnection()

	fmt.Println("Creating User model")
	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}
	db.Migrator().CreateTable(&User{})

	password := domain.NewPassword("admin")
	password.Hash()
	adminUser := &User{Username: "admin", Password: password.Password}
	db.Create(adminUser)

	// fmt.Println("Creating Manager")
	// if db.Migrator().HasTable(&domain.Manager{}) {
	// 	db.Migrator().DropTable(&domain.Manager{})
	// }
	// db.Migrator().CreateTable(&domain.Manager{})
}

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

	adminPassword := domain.NewPassword("admin")
	adminPassword.Hash()
	adminUser := &User{Username: "admin", Password: adminPassword.Password}
	db.Create(adminUser)

	managerPassword := domain.NewPassword("manager")
	managerPassword.Hash()
	managerUser := &User{Username: "manager", Password: managerPassword.Password}
	db.Create(managerUser)

	fmt.Println("Creating Admin model")
	if db.Migrator().HasTable(&Admin{}) {
		db.Migrator().DropTable(&Admin{})
	}
	db.Migrator().CreateTable(&Admin{})

	admin := &Admin{Name: "admin", UserID: adminUser.ID}
	db.Create(admin)

	fmt.Println("Creating Manager model")
	if db.Migrator().HasTable(&Manager{}) {
		db.Migrator().DropTable(&Manager{})
	}
	db.Migrator().CreateTable(&Manager{})

	manager := &Manager{Name: "manager", UserID: managerUser.ID}
	db.Create(manager)

	fmt.Println("Creating Client model")
	if db.Migrator().HasTable(&Client{}) {
		db.Migrator().DropTable(&Client{})
	}
	db.Migrator().CreateTable(&Client{})

	fmt.Println("Creating Permission model")
	if db.Migrator().HasTable(&Permission{}) {
		db.Migrator().DropTable(&Permission{})
	}
	db.Migrator().CreateTable(&Permission{})
}

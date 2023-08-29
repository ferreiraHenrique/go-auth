package sqlite

import (
	"log"

	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func Migrator() {
	db := GetConnection()

	tables := []any{&Permission{}, &User{}, &Admin{}, &Manager{}, &Client{}}

	log.Println("Tables:")
	for _, t := range tables {
		log.Printf("%T\n", t)
	}

	log.Println("Droping tables")
	db.Migrator().DropTable(tables...)

	log.Println("Creating tables")
	db.AutoMigrate(tables...)

	log.Println("Creating default admin user")
	adminPassword := domain.NewPassword("admin")
	adminPassword.Hash()
	adminUser := &User{Username: "admin", Password: adminPassword.Password}
	db.Create(adminUser)

	log.Println("Creating default admin")
	admin := &Admin{Name: "admin", UserID: adminUser.ID}
	db.Create(admin)

	log.Println("Creating default manager user")
	managerPassword := domain.NewPassword("manager")
	managerPassword.Hash()
	managerUser := &User{Username: "manager", Password: managerPassword.Password}
	db.Create(managerUser)

	log.Println("Creating default admin")
	manager := &Manager{Name: "manager", UserID: managerUser.ID}
	db.Create(manager)
}

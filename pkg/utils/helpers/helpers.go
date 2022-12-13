package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Remember to capitalize the function name,
Otherwise, the function will not be exported.
*/
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/go_fintech"), &gorm.Config{})
	HandleErr(err)
	return db
}

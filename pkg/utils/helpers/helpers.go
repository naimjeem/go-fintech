package helpers

import (
	"golang.org/x/crypto/bcrypt"
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

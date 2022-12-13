package migrations

import (
	"github.com/naimjeem/go-fintech/pkg/utils/helpers"
	"github.com/naimjeem/go-fintech/pkg/utils/interfaces"
	// "github.com/spf13/viper"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := [2]interfaces.User{
		{Username: "naimuddin", Email: "naimuddin.shahjalal@brainstation-23.com"},
		{Username: "mahtab", Email: "mahtab@brainstation-23.com"},
	}

	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	// viper.SetConfigFile("./pkg/common/envs")
	// viper.ReadInConfig()

	// port := viper.Get("PORT").(string)
	// dbUrl := viper.Get("DB_URL").(string)
	db := helpers.ConnectDB()
	db.AutoMigrate(&interfaces.User{}, &interfaces.Account{})

	createAccounts()
}

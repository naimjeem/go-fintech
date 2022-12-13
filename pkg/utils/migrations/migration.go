package migrations

import (
	"github.com/naimjeem/go-fintech/pkg/utils/helpers"
	// "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	gorm.Model
	Type    string `json:"type"`
	Name    string `json:"name"`
	Balance uint   `json:"balance"`
	UserID  uint   `json:"userId"`
}

func connectDB(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	helpers.HandleErr(err)
	return db
}

func createAccounts(url string) {
	db := connectDB(url)

	users := [2]User{
		{Username: "naimuddin", Email: "naimuddin.shahjalal@brainstation-23.com"},
		{Username: "mahtab", Email: "mahtab@brainstation-23.com"},
	}

	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	// viper.SetConfigFile("./pkg/common/envs")
	// viper.ReadInConfig()

	// port := viper.Get("PORT").(string)
	// dbUrl := viper.Get("DB_URL").(string)
	db := connectDB("postgres://postgres:root@localhost:5432/go_fintech")
	db.AutoMigrate(&User{}, &Account{})

	createAccounts("postgres://postgres:root@localhost:5432/go_fintech")
}

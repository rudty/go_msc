package helpers

import (
	"net/url"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	u := url.URL{}
	u.User = url.UserPassword("selectman", "1234")
	u.Host = "tcp(127.0.0.1:3306)"
	u.Path = "my_database"
	arg := url.Values{}
	arg.Set("loc", "Asia/Seoul")
	u.RawQuery = arg.Encode()

	db, err := gorm.Open("mysql", u.String()[2:])
	HandleErr(err)
	db.LogMode(true)
	return db
}

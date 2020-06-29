package helpers

import (
	"bank/interfaces"
	"net/url"
	"regexp"

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

func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}

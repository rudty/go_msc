package users

import (
	"bank/helpers"
	"bank/interfaces"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) (map[string]interface{}, bool) {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		// Connect DB
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("username = ? ", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}, false
		}
		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}, false
		}
		// Find accounts for the user
		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

		defer db.Close()

		var response = prepareResponse(user, accounts)

		return response, true
	} else {
		return map[string]interface{}{"message": "not valid values"}, false
	}
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func Register(username string, email string, pass string) (map[string]interface{}, bool) {
	// Add validation to registration
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := helpers.ConnectDB()
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(username + "'s" + " account"), Balance: 0, UserID: user.ID}
		db.Create(&account)

		defer db.Close()
		accounts := []interfaces.ResponseAccount{}
		respAccount := interfaces.ResponseAccount{ID: account.ID, Name: account.Name, Balance: int(account.Balance)}
		accounts = append(accounts, respAccount)
		var response = prepareResponse(user, accounts)

		return response, true
	} else {
		return map[string]interface{}{"message": "not valid values"}, false
	}
}

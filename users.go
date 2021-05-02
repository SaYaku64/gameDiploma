package main

import (
	logger "logger"

	"github.com/gin-gonic/gin"
)

// AllUsersMap - cache of all users
var AllUsersMap = AllUsers{
	Cache: make(map[uint64]User),
}

// GetUserByInfo - return user if login and password are correct
func (au *AllUsers) GetUserByInfo(login, password string) (User, bool) {
	for _, user := range au.Cache {
		if login == user.Login && checkEncoded(password, user.Password) {
			return user, true
		}
	}
	return User{}, false
}

// FillAllUsers - fills map of all users
func (au *AllUsers) FillAllUsers() {
	users := GetAllUsers()
	for _, user := range users {
		AllUsersMap.RLock()
		AllUsersMap.Cache[user.ID] = user
		AllUsersMap.RUnlock()
	}
}

// UpsertUserToMap - adds token field to map of users
func (au *AllUsers) UpsertUserToMap(user User, token string, endless bool) {
	user.Token.Token = token
	user.Token.Endless = endless
	au.RLock()
	au.Cache[user.ID] = user
	au.RUnlock()
}

// GetUserByToken - return user by token
func (au *AllUsers) GetUserByToken(token string) (User, bool) {
	for _, user := range au.Cache {
		if user.Token.Token == token {
			return user, true
		}
	}
	return User{}, false
}

// GetCurrentUser - gets user by token from gin.Context
func GetCurrentUser(c *gin.Context) (User, bool) {
	token, err := c.Cookie("nules")
	if err != nil {
		logger.Error.Println("users.go -> GetCurrentUser -> Cookie: err =", err)
		return User{}, false
	}

	return AllUsersMap.GetUserByToken(token)
}

func checkEmailExist(email string) bool {
	for _, user := range AllUsersMap.Cache {
		if email == user.Email {
			return true
		}
	}
	return false
}

func checkLoginExist(login string) bool {
	for _, user := range AllUsersMap.Cache {
		if login == user.Login {
			return true
		}
	}
	return false
}

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
		if login == user.Login && password == user.Password {
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

// GetUserByToken - return user by token
func (au *AllUsers) GetUserByToken(token string) (User, bool) {
	for _, user := range au.Cache {
		if user.Token == token {
			return user, true
		}
	}
	return User{}, false
}

// AddTokenToMap - adds token fielda to map of users
func (au *AllUsers) AddTokenToMap(id uint64, token string, endless bool) {
	au.RLock()
	user := au.Cache[id]
	user.Token = token
	user.Endless = endless
	au.Cache[id] = user
	au.RUnlock()
}

// UpdateTokenInDB - updates token in cache and DB
func (au *AllUsers) UpdateTokenInDB(user User, token string, endless bool) bool {
	if err := UpdateID(user.ID, obj{"$set": obj{"token": token, "endless": endless}}); err != nil {
		logger.Error.Println("users.go -> UpdateTokenInDB -> UpdateID: err =", err)
		return false
	}
	au.AddTokenToMap(user.ID, token, endless)
	return true
}

// GetCurrentUser - gets user by token from gin.Context
func GetCurrentUser(c *gin.Context) (User, bool) {
	token, err := c.Cookie("token")
	if err != nil {
		logger.Error.Println("users.go -> GetCurrentUser -> Cookie: err =", err)
		return User{}, false
	}

	return AllUsersMap.GetUserByToken(token)
}

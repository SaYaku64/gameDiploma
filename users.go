package main

import (
	"net/http"

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
		au.RLock()
		au.Cache[user.ID] = user
		au.RUnlock()
	}
}

// GetUserBuildings - gets current user's buildings
func GetUserBuildings(c *gin.Context) ([]DBBuilding, bool) {
	user, ok := GetCurrentUser(c)
	if !ok {
		return []DBBuilding{}, ok
	}
	return user.Fields, ok
}

// Changes territory
func changeTerritory(c *gin.Context) {
	user, ok := GetCurrentUser(c)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("unauth", Errors),
		})
		return
	}
	territory := c.PostForm("territory")
	build := c.PostForm("build")

	if len(user.Fields) == 0 {
		user.Fields = append(user.Fields, DBBuilding{TID: territory, BID: build})
		ok := AllUsersMap.UpdateBuildingsInDB(user, user.Fields)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"message": getLoc("successfulTerrChanging", Errors),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("invalidTerrChanging", Errors),
		})
		return
	}

	var changed bool

	for i, v := range user.Fields {
		if v.TID == territory {
			user.Fields[i].BID = build
			changed = true
			continue
		}
	}
	if !changed {
		user.Fields = append(user.Fields, DBBuilding{TID: territory, BID: build})
	}

	ok = AllUsersMap.UpdateBuildingsInDB(user, user.Fields)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": getLoc("successfulTerrChanging", Errors),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   true,
		"message": getLoc("invalidTerrChanging", Errors),
	})
}

// GetUserBuildingByTID - gets current user's building by territory
func GetUserBuildingByTID(c *gin.Context, tid string) string {
	dbBuildings, ok := GetUserBuildings(c)
	if !ok || len(dbBuildings) == 0 {
		return ""
	}
	for _, bldStruct := range dbBuildings {
		if tid == bldStruct.TID {
			return bldStruct.BID
		}
	}
	return ""
}

// UpdateUserInMap - updates user in map of users
func (au *AllUsers) UpdateUserInMap(user User) {
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
		Error.Println("users.go -> GetCurrentUser -> Cookie: err =", err)
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

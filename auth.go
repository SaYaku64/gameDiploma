// handlers.user.go

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quickemailverification/quickemailverification-go"
)

// Generates random token
func generateSessionToken() string {
	return fmt.Sprint(time.Now().UnixNano())
}

func checkEmailValidation(email string) bool {

	//return true ////////

	qev := quickemailverification.CreateClient(quickEmail)
	// Need to use Verify instead Sandbox in production
	resp, err := qev.Sandbox(email) // Email address which need to be verified
	if err != nil {
		Info.Printf("auth.go -> checkEmailValidation: error = %s; email = %s\n", err, email)
		return false
	}
	if resp.Result == "valid" {
		return true
	}
	return false
}

func performLogin(c *gin.Context) {
	login := c.PostForm("username")
	password := c.PostForm("password")
	check, _ := strconv.ParseBool(c.PostForm("check"))

	if strings.TrimSpace(login) != "" && strings.TrimSpace(password) != "" {
		user, ok := AllUsersMap.GetUserByInfo(login, password)

		if ok {
			c.Set("logged", true)
			token := generateSessionToken()
			var time int
			if check {
				time = 60 * 60 * 24 * 365 // cookie for 1y
				AllUsersMap.UpdateTokenInDB(user, token, true)
			} else {
				time = 60 * 60 * 2 // cookie for 2h
				AllUsersMap.UpdateTokenInDB(user, token, false)
			}
			c.SetCookie("nules", token, time, "", "", false, true)

			c.JSON(http.StatusOK, gin.H{
				"message": getLoc("successfulLogin", Errors),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("invalidLogin", Errors),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("emptyLoginField", Errors),
		})
	}
}

// logout - Deletes tokens and cookie
func logout(c *gin.Context) {
	if user, ok := GetCurrentUser(c); ok {
		if ok := AllUsersMap.UpdateTokenInDB(user, "", false); ok {
			c.SetCookie("nules", "", -1, "", "", false, true)
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	}
}

// Adds new user to DB
func register(c *gin.Context) {
	email := c.PostForm("email")

	ok := checkEmailValidation(email)
	if ok {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if token, err := registerNewUser(email, username, password); err == "" {
			time := 60 * 60 * 2
			c.SetCookie("nules", token, time, "", "", false, true) // token 10m
			c.Set("logged", true)

			c.JSON(http.StatusOK, gin.H{
				"message": getLoc("successfulRegistration", Errors),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error":   true,
				"message": err,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": getLoc("invalidEmail", Errors),
		})
	}
}

func registerNewUser(email, username, password string) (string, string) {
	if strings.TrimSpace(password) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(username) == "" {
		return "", getLoc("emptyRegField", Errors)
	} else if checkEmailExist(email) {
		return "", getLoc("notUniqueEmail", Errors)
	} else if checkLoginExist(username) {
		return "", getLoc("notUniqueLogin", Errors)
	}

	hPass := encode(password)
	userToken := generateSessionToken()
	usr := User{ID: uint64(len(AllUsersMap.Cache) + 1), Login: username, Email: email, Password: hPass, Token: token{userToken, false}}

	if err := AllUsersMap.AddNewUser(usr); err != nil {
		Error.Println("auth.go -> registerNewUser -> AddNewUser: err =", err)
		return "", getLoc("internalError", Errors)
	}

	return userToken, ""
}

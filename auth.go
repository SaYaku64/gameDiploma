// handlers.user.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quickemailverification/quickemailverification-go"
)

// Generates random token
func generateSessionToken() string {
	// hash, _ := HashString("AcCeSs noT deNIeD")
	// return hash
	return fmt.Sprint(time.Now().UnixNano())
}

func checkEmailValidation(email string) string {
	qev := quickemailverification.CreateClient(quickEmail)
	// Need to use Verify instead Sandbox in production
	response, err := qev.Sandbox(email) // Email address which need to be verified
	if err != nil {
		log.Println(err)
		return "Validation failed"
	}
	return response.Result
}

func performLogin(c *gin.Context) {
	login := c.PostForm("usernameLogin")
	password := c.PostForm("passwordLogin")
	check := c.PostForm("checkLogin")

	if strings.TrimSpace(login) != "" && strings.TrimSpace(password) != "" {
		user, ok := AllUsersMap.GetUserByInfo(login, password)

		if ok {
			c.Set("logged", true)
			token := generateSessionToken()
			var time int
			if check == "true" {
				time = 60 * 60 * 24 * 365 // cookie for 1y
				AllUsersMap.UpdateTokenInDB(user, token, true)
			} else {
				time = 60 * 60 * 2 // cookie for 2h
				AllUsersMap.UpdateTokenInDB(user, token, false)
			}
			c.SetCookie("token", token, time, "", "", false, true)

			c.JSON(200, gin.H{
				"message": "Successful login",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"error":   true,
			"message": "Login Failed: Invalid login or password",
		})
	}
}

// Logout - Deletes tokens and cookie
func Logout(c *gin.Context) {
	user, ok := GetCurrentUser(c)
	if ok {
		if ok := AllUsersMap.UpdateTokenInDB(user, "", false); ok {
			c.SetCookie("token", "", -1, "", "", false, true)
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	}
}

// Adds new user to DB
// func register(c *gin.Context) {
// 	email := c.PostForm("email")

// 	result := checkEmailValidation(email)
// 	if result == "valid" {
// 		username := c.PostForm("username")
// 		password := c.PostForm("password")

// 		if token, err := registerNewUser(email, username, password); err == nil {
// 			c.SetCookie("token", token, 600, "", "", false, true) // token 10m
// 			c.Set("logged", true)

// 			// showIndexPage(c)
// 			render(c, gin.H{
// 				"title": "Home page"}, "index.html")
// 		} else {
// 			render(c, gin.H{
// 				"title":        "Register",
// 				"ErrorTitle":   "Registration Failed",
// 				"ErrorMessage": err.Error(),
// 			}, "register.html")
// 		}
// 	} else {
// 		render(c, gin.H{
// 			"title":        "Register",
// 			"ErrorTitle":   "Registration Failed",
// 			"ErrorMessage": "Invalid email adress",
// 		}, "register.html")
// 	}

// }

// func registerNewUser(email, username, password string) (string, error) {
// 	if strings.TrimSpace(password) == "" {
// 		return "", errors.New("The password field can't be empty")
// 	} else if strings.TrimSpace(email) == "" {
// 		return "", errors.New("The email adress field can't be empty")
// 	} else if !checkEmailExist(email) {
// 		return "", errors.New("The email is already used")
// 	} else if strings.TrimSpace(username) == "" {
// 		return "", errors.New("The username field can't be empty")
// 	} else if !checkUserExist(username) {
// 		return "", errors.New("The username is already used")
// 	}

// 	hPass, err := HashString(password)
// 	if err != nil {
// 		return "", err
// 	}
// 	u := user{Email: email, Username: username, Password: hPass, Token: token{Name: generateSessionToken()}}

// 	err = addUserToDB(u)
// 	if err != nil {
// 		return "", err
// 	}

// 	u.addUserToCache()

// 	return u.Token.Name, nil
// }

// middleware.auth.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var ginErrors = map[string]ginError{
// 	"notLogged": {,},
// 	"logged":    {"Auth error", "You are already signed in!"},
// 	"unknown":   {"Unexpected error", "Please continue using site later."},
// }

// IsLoggedIn - checks if user has "logged" status
func IsLoggedIn(c *gin.Context) bool {
	loggedInInterface, ok := c.Get("logged")
	if ok && loggedInInterface.(bool) {
		return true
	}
	return false
}

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsLoggedIn(c) {
			c.AbortWithStatus(http.StatusUnauthorized)
			render(c, "index.html", gin.H{
				"title":        "NULESandbox",
				"errorTitle":   "Auth error",
				"errorMessage": "Please sign in to visit this page!",
			})
		}

	}
}

// func ensureNotLoggedIn() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		loggedInInterface, ok := c.Get("logged")
// 		if ok {
// 			loggedIn := loggedInInterface.(bool)
// 			if loggedIn {
// 				c.AbortWithStatus(http.StatusUnauthorized)
// 				render(c, gin.H{
// 					"title": "Home Page",
// 				}, "index.html")
// 			}
// 		} else {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			render(c, gin.H{
// 				"title": "Home Page",
// 			}, "index.html")
// 		}

// 	}
// }

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		Localization = getLocalization(c)

		if token, err := c.Cookie("token"); err == nil || token != "" { // take cookie

			user, found := AllUsersMap.GetUserByToken(token)

			if found {
				var time int
				if !user.Endless {
					time = 60 * 60 * 2 // cookie for 2h
				} else {
					time = 60 * 60 * 24 * 365 // cookie for 1y
				}
				c.SetCookie("token", token, time, "", "", false, true)
				c.Set("logged", true)
			}

		} else {
			c.Set("logged", false)
		}
	}
}

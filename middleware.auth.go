// middleware.auth.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsLoggedIn - checks if user has "logged" status
func IsLoggedIn(c *gin.Context) bool {
	loggedInInterface, ok := c.Get("logged")
	if ok && loggedInInterface.(bool) {
		return true
	}
	return false
}

// IsAsked - checks if user has "asked" status
func IsAsked(c *gin.Context) bool {
	askedInterface, ok := c.Get("asked")
	if ok && askedInterface.(bool) {
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
				"messageError": getLoc("unauth", Errors),
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

		if token, err := c.Cookie("nules"); err == nil || token != "" { // take cookie

			user, found := AllUsersMap.GetUserByToken(token)

			if found {
				var time int
				if !user.Token.Endless {
					time = 60 * 60 * 2 // cookie for 2h
				} else {
					time = 60 * 60 * 24 * 365 // cookie for 1y
				}
				c.SetCookie("nules", token, time, "", "", false, true)
				c.Set("asked", user.Asked)
				c.Set("logged", true)
			}

		} else {
			c.Set("asked", false)
			c.Set("logged", false)
		}
	}
}

package main

////////////////////////////////////////////
// Checking admin status
////////////////////////////////////////////
// func ensureAdminned() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		adminnedInterface, _ := c.Get("adminned")
// 		adminnedIn := adminnedInterface.(bool)

// 		if !adminnedIn {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			render(c, gin.H{
// 				"title": "Home Page",
// 			}, "index.html")
// 		}
// 	}
// }

// func setAdminStatus() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if spToken, err := c.Cookie("specialToken"); err == nil && spToken != "" { // take cookie

// 			if token, err := c.Cookie("token"); err == nil && token != "" { // take cookie
// 				ActiveUsers.RLock()
// 				if u, ok := ActiveUsers.m[token]; ok { // take user from cache
// 					ActiveUsers.RUnlock()
// 					if u.AdminStatus && spToken == u.AdminToken {
// 						c.Set("adminned", true)
// 					} else {
// 						c.SetCookie("token", token, 300, "", "", false, true) // token 5m
// 						c.Set("adminned", true)
// 					}
// 				}
// 			}
// 		} else {
// 			c.Set("adminned", false)
// 		}
// 	}
// }

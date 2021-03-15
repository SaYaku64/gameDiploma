// routes.go

package main

func initializeRoutes() {

	// Use the setUserStatus and setAdminStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())
	//router.Use(setAdminStatus())

	// Handling Index-page
	router.GET("/", showIndexPage)
	//router.POST("/", performLogin)

	router.POST("/coord", Coords)

	// adminRoutes := router.Group("/admin")
	// {
	// 	adminRoutes.GET("/panel", ensureAdminned(), showAdminPanelPage)

	// 	adminRoutes.POST("/panel", ensureAdminned(), delThisShit)
	// }

}

// routes.go

package main

func initializeRoutes() {

	// Use the setUserStatus and setAdminStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handling Index-page
	router.GET("/", showIndexPage)
	router.POST("/login", performLogin)
	router.POST("/register", register)
	router.POST("/change", changeTerritory)
	router.GET("/logout", logout)
	router.GET("/survey", ensureLoggedIn(), showSurveyPage)
	router.POST("/survey", ensureLoggedIn(), surveyComplete)
	router.GET("/infographics", ensureLoggedIn(), showInfographicsPage)

	router.GET("/object/:name", showObjectPage)

	//router.POST("/", performLogin)

	// router.POST("/mapSize", MapSize)
	// router.POST("/coord", Coords)

	// adminRoutes := router.Group("/admin")
	// {
	// 	adminRoutes.GET("/panel", ensureAdminned(), showAdminPanelPage)

	// 	adminRoutes.POST("/panel", ensureAdminned(), delThisShit)
	// }

}

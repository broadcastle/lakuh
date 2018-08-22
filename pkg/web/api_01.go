package web

func api1() {

	a := e.Group("/api/v1")

	// User

	u := a.Group("/user")

	u.POST("/", userCreate)
	u.POST("/login/", userLogin)

	pr := u.Group("/profile")

	useJWT(pr)

	pr.GET("/", userRead)
	pr.PUT("/", userUpdate)
	pr.DELETE("/", userDelete)

	// Library

	l := a.Group("/library")

	useJWT(l)

	l.POST("/", libraryAdd)
	l.GET("/", libraryView)

	li := l.Group("/:audio")

	li.GET("/", songView)
	li.PUT("/", songUpdate)
	li.DELETE("/", songDelete)

	// Player

}

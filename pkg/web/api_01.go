package web

func api1() {

	a := e.Group("/api/v1")

	// USER

	u := a.Group("/user")

	u.POST("/", userCreate)
	u.POST("/login/", userLogin)

	p := u.Group("/profile")

	useJWT(p)

	p.GET("/", userRead)
	p.PUT("/", userUpdate)
	p.DELETE("/", userDelete)

}

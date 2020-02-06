package app

import (
		"bookstore_users-api/controllers/ping"
	users "bookstore_users-api/controllers/user"
)

func mapUrls(){
    router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.Delete)
}

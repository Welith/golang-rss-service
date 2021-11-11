package main

import "github.com/gin-gonic/gin"

func main()  {

	r := gin.Default()

	r.POST("/login", Login)
	r.POST("/v1/feeds", ParseFeed)

	err := r.Run(":3000")

	if err != nil {

		LogError(err.Error())
		panic(err)
	}
}
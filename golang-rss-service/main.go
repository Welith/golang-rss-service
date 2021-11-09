package main

import "github.com/gin-gonic/gin"

func main()  {

	r := gin.Default()
	r.POST("/v1/feeds", ParseFeed)

	err := r.Run(":8080")

	if err != nil {

		LogError(err.Error())
		panic(err)
	}
}
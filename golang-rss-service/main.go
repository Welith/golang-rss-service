package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"path/filepath"
)

func main()  {

	err := godotenv.Load(filepath.Join(".", ".env"))

	if err != nil {

		panic(err)
	}

	//panic("test")
	go ProcessIncomingMessages()

	r := gin.Default()

	r.POST("/login", Login)
	r.POST("/logout", TokenAuthMiddleware(), Logout)
	r.POST("/v1/feeds", TokenAuthMiddleware(), ParseFeed)
	r.POST("/token/refresh", Refresh)

	if err := r.Run(":3000"); err != nil {

		LogError(err.Error())
		panic(err)
	}
}
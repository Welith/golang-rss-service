package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func LogError(exception string) {

	// Enable logging
	var filename string

	if os.Getenv("ENV") == "dev" {

		filename, _ = filepath.Abs("logs/prod.log") // TODO: change to ../logs
	} else {

		filename, _ = filepath.Abs("../logs/prod.log") // TODO: change to ../logs
	}

	if !LogFileExists(filename) {

		err := CreateLogFile(filename)
		if err != nil {

			panic(err)
		}
	}

	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	defer f.Close()

	log.SetOutput(f)

	log.Println("-------------------------ERROR-------------------------")
	log.Println(exception)
	log.Println("-------------------------ERROR-------------------------")
}

func ErrorResponse(c *gin.Context, exception *ErrorResponseStruct) {

	LogError(exception.ErrorMsg)
	c.JSON(http.StatusOK, exception)
}


package main

import (
	golang_rss_reader_package "github.com/Welith/golang-rss-reader-package"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseFeed(c *gin.Context)  {

	request := new(Request)
	err := c.BindJSON(&request)

	if err != nil {

		exception := BuildErrorResponse(StatusText(HttpBadRequest), err.Error())
		ErrorResponse(c, exception)
		return
	}

	err = request.Validate()

	if err != nil {

		exception := BuildErrorResponse(StatusText(ValidationError), err.Error())
		ErrorResponse(c, exception)
		return
	}

	result := golang_rss_reader_package.Parse(request.Urls)

	var response ResponseItem

	response.Items = result

	c.JSON(http.StatusOK, response)
}

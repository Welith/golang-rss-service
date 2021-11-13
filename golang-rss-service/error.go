package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	General               = 11000
	ValidationError       = 11001
	HttpBadRequest        = 400
	StructNotFound        = 11002
	HttpUnauthorized      = 401
	HttpNotFoundException = 404
	HttpServerException   = 500
	NonUniqueValue        = 11004
	ServiceException      = 11008
)

var statusText = map[int]string{

	General:               "services.general.bad_request",
	ValidationError:       "service.general.validation_error",
	HttpBadRequest:        "service.general.bad_request",
	StructNotFound:        "service.general.struct_not_found",
	HttpUnauthorized:      "service.general.unauthorized",
	HttpNotFoundException: "service.general.http_not_found",
	HttpServerException:   "service.general.service_request_exception",
	NonUniqueValue:        "service.general.non_unique_value",
	ServiceException:      "service.general.service_exception",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

//LogError proper error logging
func LogError(exception string) {

	filename, _ := filepath.Abs("logger.log") // TODO: change to ../logs

	if !LogFileExists(filename) {

		if err := CreateLogFile(filename); err != nil {

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
	c.JSON(http.StatusBadRequest, exception)
}


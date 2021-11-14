package main

import (
	"fmt"
	golang_rss_reader_package "github.com/Welith/golang-rss-reader-package"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

//ParseFeed action endpoint
func ParseFeed(c *gin.Context)  {

	tokenAuth, err := ExtractTokenMetadata(c.Request)

	if err != nil {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), err.Error())
		ErrorResponse(c, exception)
		return
	}
	_, err = FetchAuth(tokenAuth)

	if err != nil {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), err.Error())
		ErrorResponse(c, exception)
		return
	}

	request := new(Request)

	if err = c.BindJSON(&request); err != nil {

		exception := BuildErrorResponse(StatusText(HttpBadRequest), err.Error())
		ErrorResponse(c, exception)
		return
	}

	if err = request.Validate(); err != nil {

		exception := BuildErrorResponse(StatusText(ValidationError), err.Error())
		ErrorResponse(c, exception)
		return
	}

	result := golang_rss_reader_package.Parse(request.Urls)

	var response ResponseItem

	response.Items = result

	if response.Items == nil {

		response.Items = []golang_rss_reader_package.RssItem{}
	}

	c.JSON(http.StatusOK, response)
}

//Login endpoint (JWT)
func Login(c *gin.Context) {

	var u User

	if err := c.ShouldBindJSON(&u); err != nil {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), err.Error())
		ErrorResponse(c, exception)
		return
	}

	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), "Please provide valid login details")
		ErrorResponse(c, exception)
		return
	}

	ts, err := CreateToken(user.ID)

	if err != nil {

		exception := BuildErrorResponse(StatusText(HttpBadRequest), "Please provide valid login details")
		ErrorResponse(c, exception)
		return
	}

	saveErr := CreateAuth(user.ID, ts)

	if saveErr != nil {

		exception := BuildErrorResponse(StatusText(HttpBadRequest), saveErr.Error())
		ErrorResponse(c, exception)
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

//Logout API (JWT)
func Logout(c *gin.Context) {

	au, err := ExtractTokenMetadata(c.Request)

	if err != nil {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), err.Error())
		ErrorResponse(c, exception)
		return
	}

	deleted, delErr := DeleteAuth(au.AccessUuid)

	if delErr != nil || deleted == 0 { //if any goes wrong

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), delErr.Error())
		ErrorResponse(c, exception)
		return
	}

	c.JSON(http.StatusOK, "Successfully logged out")
}


func Refresh(c *gin.Context) {

	mapToken := map[string]string{}

	if err := c.ShouldBindJSON(&mapToken); err != nil {

		exception := BuildErrorResponse(StatusText(HttpBadRequest), err.Error())
		ErrorResponse(c, exception)
		return
	}

	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {

		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), err.Error())
		ErrorResponse(c, exception)
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {

		exception := BuildErrorResponse(StatusText(HttpUnauthorized), "Token invalid!")
		ErrorResponse(c, exception)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

	if ok && token.Valid {

		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string

		if !ok {

			exception := BuildErrorResponse(StatusText(HttpBadRequest), "Error")
			ErrorResponse(c, exception)
			return
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {

			exception := BuildErrorResponse(StatusText(HttpBadRequest), err.Error())
			ErrorResponse(c, exception)
			return
		}

		//Delete the previous Refresh Token
		deleted, delErr := DeleteAuth(refreshUuid)

		if delErr != nil || deleted == 0 { //if any goes wrong

			exception := BuildErrorResponse(StatusText(HttpUnauthorized), delErr.Error())
			ErrorResponse(c, exception)
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := CreateToken(userId)

		if  createErr != nil {

			exception := BuildErrorResponse(StatusText(http.StatusForbidden), createErr.Error())
			ErrorResponse(c, exception)
			return
		}
		//save the tokens metadata to redis
		saveErr := CreateAuth(userId, ts)

		if saveErr != nil {

			exception := BuildErrorResponse(StatusText(http.StatusForbidden), saveErr.Error())
			ErrorResponse(c, exception)
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		c.JSON(http.StatusCreated, tokens)
	} else {

		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
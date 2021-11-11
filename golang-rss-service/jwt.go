package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TokenDetails struct {

	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type User struct {

	ID 		 uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
//A sample use
var user = User{

	ID:        1,
	Username: "emerchantpay",
	Password: "password",
}

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
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

func CreateToken(userid uint64) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {

		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	if err != nil {

		return nil, err
	}
	return td, nil
}

func CreateAuth(userid uint64, td *TokenDetails) error {

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()

	if errAccess != nil {

		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()

	if errRefresh != nil {

		return errRefresh
	}
	return nil
}
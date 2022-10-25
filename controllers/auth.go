package controllers

import (
	"errors"
	"fmt"
	api "myapp/api"
	"myapp/config"
	"myapp/db"
	"myapp/libs"
	"myapp/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthSignUp(c echo.Context) error {
	fmt.Println("==================SignUp==================")
	db := db.DbManager()

	params := signupParameters{}
	e := c.Bind(&params)
	if e != nil {
		return api.SendError(c, 403, "bad signup params")

	}

	user := &models.User{}
	if params.Username == "" || params.Password == "" {
		return api.SendError(c, 403, "bad params")
	}

	trans := db.First(&user, "\"username\"=?", params.Username)
	if trans.Error == nil {
		return api.SendError(c, 405, "duplicate user")

	}

	user.Username = params.Username
	user.Password = params.Password
	trans = db.Save(&user)
	if trans.Error != nil {
		return api.SendError(c, 500, trans.Error.Error())

	}
	return api.SendResponse(c, 200, map[string]interface{}{
		"user": user,
	})
}

func AuthLogin(c echo.Context) error {
	fmt.Println("==================Login==================")
	db := db.DbManager()
	params := loginParameters{}
	e := c.Bind(&params)

	if e != nil {
		return api.SendError(c, 403, "bad login params")

	}
	if params.Username == "" || params.Password == "" {
		return api.SendError(c, 403, "bad params")

	}
	user := &models.User{}
	trans := db.First(&user, "username = ?", params.Username)
	if trans.Error != nil {
		if errors.Is(trans.Error, gorm.ErrRecordNotFound) {
			return api.SendError(c, 404, "user not found")

		}
		return api.SendError(c, 500, trans.Error.Error())

	}
	if user.Password != params.Password {
		return api.SendError(c, 405, "wrong password")

	}
	//generate access token:
	newToken, err := libs.JWTEncode(config.JWT_SECRET, jwt.MapClaims{
		"id": user.ID,
	}, config.JWT_TOKEN_LIFETIME)
	if err != nil {
		return api.SendError(c, 500, err.Error())

	}
	user.Token = newToken

	trans = db.Save(&user)
	if trans.Error != nil {
		return api.SendError(c, 500, trans.Error.Error())

	}

	return api.SendResponse(c, 200, map[string]interface{}{
		"user": user,
	})

}
func AuthCheckToken(c echo.Context) error {
	fmt.Println("==================CheckToken==================")
	db := db.DbManager()

	params := checktokenParameters{}
	e := c.Bind(&params)
	if e != nil {
		return api.SendError(c, 403, "bad checkToken params")
	}

	if params.Token == "" {
		return api.SendError(c, 403, "bad params")
	}

	_, err := libs.JWTDecode(config.JWT_SECRET, params.Token)
	if err != nil {
		return api.SendError(c, 403, err.Error())
	}
	user := &models.User{}
	trans := db.First(&user, "token = ?", params.Token)
	if trans.Error != nil {
		fmt.Print(trans.Error.Error())
		// if errors.Is(trans.Error, gorm.ErrRecordNotFound) {
		// 	return api.SendError(c, 404, "user not found")

		// }
		// return api.SendError(c, 500, trans.Error.Error())
		return api.SendError(c, 405, "invalidToken")

	}

	return api.SendResponse(c, 200, map[string]interface{}{
		"user": user,
	})

}

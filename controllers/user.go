package controllers

import (
	"fmt"
	api "myapp/api"

	"myapp/db"

	"myapp/models"

	"github.com/labstack/echo/v4"
)

func UserGetAllUsers(c echo.Context) error {
	fmt.Println("==================GetUsers==================")
	db := db.DbManager()

	users := []models.User{}
	trans := db.Find(&users)
	if trans.Error != nil {
		fmt.Print(trans.Error.Error())
		return api.SendError(c, 405, trans.Error.Error())

	}

	return api.SendResponse(c, 200, map[string]interface{}{
		"users": users,
	})

}

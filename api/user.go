package api

import (
	"net/http"

	"music-library-management/models"
	"music-library-management/sdk/common"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

// Handler for registering a new user
func register(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	hashedPassword, _ := common.HashPassword(u.Password)
	u.Password = hashedPassword

	collection := GetCollection(Client, "users")
	_, err := collection.InsertOne(MongoCtx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, u)
}

// Handler for logging in a user
func login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	collection := GetCollection(Client, "users")
	var result User
	err := collection.FindOne(MongoCtx, bson.M{"username": u.Username}).Decode(&result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if common.CheckPasswordHash(u.Password, result.Password) {
		return c.String(http.StatusOK, "Login successful")
	} else {
		return c.String(http.StatusUnauthorized, "Login failed")
	}
}

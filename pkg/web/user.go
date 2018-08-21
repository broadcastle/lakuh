package web

import (
	"fmt"
	"net/http"
	"time"

	"broadcastle.co/code/lakuh/pkg/db"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func userCreate(c echo.Context) error {

	user := db.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := user.Create(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)

}

func userRead(c echo.Context) error {

	user := db.User{}

	i := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64)

	user.ID = uint(i)

	if err := user.Find(); err != nil {
		return c.JSON(http.StatusMethodNotAllowed, err)
	}

	return c.JSON(http.StatusOK, user)
}

func userUpdate(c echo.Context) error {

	user := db.User{}

	i := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user.ID = uint(i)

	if err := user.Update(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func userDelete(c echo.Context) error {

	user := db.User{}

	i := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64)

	user.ID = uint(i)

	if err := user.Delete(); err != nil {
		return c.JSON(http.StatusMethodNotAllowed, err)
	}

	return c.JSON(http.StatusOK, "success")
}

func userLogin(c echo.Context) error {

	user := db.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println(user)

	if err := user.Login(); err != nil {
		return c.JSON(http.StatusMethodNotAllowed, err)
	}

	fmt.Println(user)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(viper.GetString("lakuh.token")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, t)

}
